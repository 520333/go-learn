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

// ServerStreamingEcho 服务端流式处理RPC方式实现
func (s *server) ServerStreamingEcho(req *proto.EchoRequest, stream proto.Echo_ServerStreamingEchoServer) error {
	fmt.Println("------------ StreamingEcho Server ------------")

	for i := 0; i < 5; i++ {
		err := stream.Send(&proto.EchoResponse{Message: req.Message})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamingEcho(req grpc.ClientStreamingServer[proto.EchoRequest, proto.EchoResponse]) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) BidirectionalStreamingEcho(req grpc.ClientStreamingServer[proto.EchoRequest, proto.EchoResponse]) error {
	//TODO implement me
	panic("implement me")
}
