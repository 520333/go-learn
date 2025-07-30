package main

//rpc 4种模式
//1.简单模式 Simple RPC
//2.服务端数据流模式 Server-side streaming RPC
//3.客户端数据流模式 Server-side streaming RPC
//4.双向数据流模式 Bidirectional streaming RPC
import (
	"context"
	"fmt"
	"learn/grpc/metadata/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice, "========")
		for i, e := range nameSlice {
			fmt.Print(i, e)
		}

	}
	for key, val := range md {
		fmt.Println(key, val)
	}

	return &proto.HelloReply{
		Messgae: "hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
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
