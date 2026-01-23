package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/base/proxy/grpc_proxy/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var port = flag.Int("port", 8005, "the port to server on")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterEchoServer(s, &server{})
	s.Serve(listener)
}

type server struct{}

// UnaryEcho 一元RPC服务方式实现
func (s *server) UnaryEcho(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	fmt.Println("------------ UnaryEcho Server ------------")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("miss metadata from request")
	}
	fmt.Println("md:", md)

	return &proto.EchoResponse{Message: req.Message}, nil
}

func (s *server) ServerStreamingEcho(request *proto.EchoRequest, g grpc.ServerStreamingServer[proto.EchoResponse]) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) ClientStreamingEcho(g grpc.ClientStreamingServer[proto.EchoRequest, proto.EchoResponse]) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) BidirectionalStreamingEcho(g grpc.ClientStreamingServer[proto.EchoRequest, proto.EchoResponse]) error {
	//TODO implement me
	panic("implement me")
}
