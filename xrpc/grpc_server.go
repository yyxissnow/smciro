package xrpc

import (
	"log"
	"net"
	"smicro/xrpc/interceptor"
	"time"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	name              string
	addr              string
	register          RegisterFunc
	server            *grpc.Server
	cfg               *GrpcServerConfig
	options           []grpc.ServerOption
	unaryInterceptors []grpc.UnaryServerInterceptor
}

type GrpcServerConfig struct {
	// 超时时间 默认500ms
	Timeout time.Duration
}

type RegisterFunc func(*grpc.Server)

func NewGrpcServer(name, addr string, cfg *GrpcServerConfig) *GrpcServer {
	server := &GrpcServer{
		name: name,
		addr: addr,
		cfg:  cfg,
	}
	if server.cfg == nil {
		server.cfg = &GrpcServerConfig{Timeout: time.Millisecond * 500}
	}
	return server
}

func (s *GrpcServer) RegisterService(register RegisterFunc) *GrpcServer {
	s.register = register
	if s.cfg.Timeout > 0 {
		s.AddUnaryInterceptors(interceptor.UnaryTimeout(s.cfg.Timeout))
	}
	s.AddUnaryInterceptors(interceptor.UnaryCrash1, interceptor.UnaryCrash2)
	options := append(s.options, interceptor.WithUnaryServer(s.unaryInterceptors...))
	s.server = grpc.NewServer(options...)
	return s
}

func (s *GrpcServer) Start() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("net.Listen(%s) err:%v", s.addr, err)
	}
	if s.server != nil {
		s.register(s.server)
	}
	if err = s.server.Serve(lis); err != nil {
		log.Fatalf("s.server.Serve() err:%v", err)
	}
}

func (s *GrpcServer) Close() {
	if s.server != nil {
		s.server.Stop()
	}
}
