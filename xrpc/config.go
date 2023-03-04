package xrpc

import "google.golang.org/grpc"

func (s *GrpcServer) AddUnaryInterceptors(interceptor ...grpc.UnaryServerInterceptor) {
	s.unaryInterceptors = append(s.unaryInterceptors, interceptor...)
}
