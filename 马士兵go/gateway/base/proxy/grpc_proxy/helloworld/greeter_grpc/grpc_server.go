package main

import (
	"context"
	"fmt"
	"gateway/base/proxy/grpc_proxy/helloworld/greeter_grpc/pb"
	"net"

	"google.golang.org/grpc"
)

type HelloService struct{}

func (hs *HelloService) Hello(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	reply := &pb.Person{
		Name: person.Name,
		Age:  person.Age,
	}
	return reply, nil

	/*	if err != nil {
		return pb.Hello(ctx, person)
	}*/
}

func main() {
	// 1.注册gRPC服务
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})
	listener, err := net.Listen("tcp", ":8004")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	grpcServer.Serve(listener)
}
