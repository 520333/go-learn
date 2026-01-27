package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/base/proxy/grpc_proxy/proto"
	"io"
	"log"
	"net"
	"strconv"

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
		err := stream.Send(&proto.EchoResponse{Message: req.Message + " " + strconv.Itoa(i+1)})
		if err != nil {
			return err
		}
	}
	return nil
}

// ClientStreamingEcho 客户端流式处理RPC方法实现
func (s *server) ClientStreamingEcho(stream proto.Echo_ClientStreamingEchoServer) error {
	fmt.Println("------------ ClientStreamingEcho Client ------------")
	var message = "received over!"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("echo lasted received message")
			return stream.SendAndClose(&proto.EchoResponse{Message: message})
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received: %s\n", req.Message)
	}
}

// BidirectionalStreamingEcho 双向流处理RPC方法实现
func (s *server) BidirectionalStreamingEcho(stream proto.Echo_BidirectionalStreamingEchoServer) error {
	fmt.Println("------------ BidirectionalStreamingEcho Server------------")
	var message = "received over!"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("echo lasted received message")
			return stream.SendMsg(&proto.EchoResponse{Message: message})
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received: %s\n", req.Message)
		if err = stream.SendMsg(&proto.EchoResponse{Message: "request received: " + req.Message}); err != nil {
			return err
		}
	}
}
