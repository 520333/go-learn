package main

//rpc 4种模式
//1.简单模式 Simple RPC
//2.服务端数据流模式 Server-side streaming RPC
//3.客户端数据流模式 Server-side streaming RPC
//4.双向数据流模式 Bidirectional streaming RPC
import (
	"context"
	"fmt"
	"learn/grpc/interpretor/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	return &proto.HelloReply{
		Messgae: "hello " + request.Name,
	}, nil
}

// 可以定义在这里
// myfunc Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	fmt.Println("接收到一个新的请求")
// 	return handler(ctx, req)
// }

func main() {
	// 定义一个server拦截器
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到一个新的请求")
		res, err := handler(ctx, req)
		fmt.Println("请求完成", res, err)
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
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
