package main

//rpc 4种模式
//1.简单模式 Simple RPC
//2.服务端数据流模式 Server-side streaming RPC
//3.客户端数据流模式 Server-side streaming RPC
//4.双向数据流模式 Bidirectional streaming RPC
import (
	"context"
	"learn/grpc/grpc_validate/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id:     request.Id,
		Email:  request.Email,
		Mobile: request.Mobile,
	}, nil
}

type Validate interface {
	Validate() error
}

func main() {
	var interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(Validate); ok {
			if err := r.Validate(); err != nil {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
