package interceptor

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/yyxissnow/smicro/app/log/xlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WithUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(chainUnaryServerInterceptors(interceptors...))
}

func UnaryTimeout(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		done := make(chan error, 1)
		h := func() {
			resp, err = handler(ctx, req)
			if err != nil {
				done <- err
				return
			}
			done <- nil
		}
		go h()
		select {
		case err = <-done:
			if err != nil {
				return nil, err
			}
			return resp, nil
		case <-ctx.Done():
			return nil, fmt.Errorf("%s request timeout", info.FullMethod)
		}
	}
}

func UnaryCrash1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("crash1")
	defer func() {
		if r := recover(); r != nil {
			err = toPanicError(r)
		}
	}()
	return handler(ctx, req)
}

func UnaryCrash2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("crash2")
	defer func() {
		if r := recover(); r != nil {
			err = toPanicError(r)
		}
	}()
	return handler(ctx, req)
}

func toPanicError(r interface{}) error {
	var buf [2 << 10]byte
	xlog.Errorf("[server-panic] - %v - %s", r, string(buf[:runtime.Stack(buf[:], false)]))
	return status.Errorf(codes.Internal, "panic: %v", r)
}

func chainUnaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	switch len(interceptors) {
	case 0:
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
			interface{}, error) {
			return handler(ctx, req)
		}
	case 1:
		return interceptors[0]
	default:
		last := len(interceptors) - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
			interface{}, error) {
			var chainHandler grpc.UnaryHandler
			var current int

			chainHandler = func(curCtx context.Context, curReq interface{}) (interface{}, error) {
				if current == last {
					return handler(curCtx, curReq)
				}

				current++
				resp, err := interceptors[current](curCtx, curReq, info, chainHandler)

				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}
}
