package main

import (
	"context"
	"fmt"
	"gateway/base/proxy/grpc_proxy/proto"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var msg = "this is client"

func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:8085",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.NewClient err: %v", err)
	}
	defer conn.Close()
	c := proto.NewEchoClient(conn)

	// 调用一元RPC方法
	UnaryEchoWithMedata(c, msg)
	time.Sleep(time.Second)

	// 调用服务端流式处理RPC方法
	ServerStreamingEchoWithMedata(c, msg)
	time.Sleep(time.Second)

	// 调用客户端流式处理RPC方法
	ClientStreamingEchoWithMedata(c, msg)
	time.Sleep(time.Second)

	// 调用双向流式RPC方法
	bidirectionalStreamingEchoWithMedata(c, msg)
	time.Sleep(time.Second)

}

func UnaryEchoWithMedata(c proto.EchoClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("------------ UnaryEcho Client ------------")
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	//md.Append("authorization", "token....")
	ctx = metadata.NewOutgoingContext(ctx, md)
	reply, err := c.UnaryEcho(ctx, &proto.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("failed to call UnaryEcho method error: %v", err)
	} else {
		fmt.Println(reply.Message)
	}
}

func ServerStreamingEchoWithMedata(c proto.EchoClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("------------ ServerStreaming Client ------------")
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := c.ServerStreamingEcho(ctx, &proto.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("failed to call ServerStreaming method error: %v", err)
	}
	var rpcError error
	for {
		resp, err := stream.Recv()
		if err != nil {
			rpcError = err
			break
		}
		fmt.Printf("response is :%s\n", resp.Message)
	}
	if rpcError != io.EOF {
		log.Fatalf("failed to finish streaming: %v", rpcError)
	}
}

func ClientStreamingEchoWithMedata(c proto.EchoClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("------------ ClientStreaming Client ------------")
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := c.ClientStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call ClientStreaming method error: %v", err)
	}

	for i := 0; i < 5; i++ {
		if err = stream.Send(&proto.EchoRequest{Message: msg + " " + strconv.Itoa(i+1)}); err != nil {
			log.Fatalf("failed to send error: %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to finish clientStreaming: %v", err)
	}
	fmt.Printf("response :%s\n", resp.Message)
}

func bidirectionalStreamingEchoWithMedata(c proto.EchoClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("------------ BidirectionalStreaming Client ------------")
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := c.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call BidirectionalStreaming method error: %v", err)
	}

	// 协程-发送消息
	go func() {
		for i := 0; i < 5; i++ {
			if err = stream.Send(&proto.EchoRequest{Message: msg + " " + strconv.Itoa(i+1)}); err != nil {
				log.Fatalf("Failed to send error: %v", err)
			}
		}
		_ = stream.CloseSend()
	}()

	// 协程-接收消息
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Failed to Finished BidirectionalStreaming %v", err.Error())
			break
		}
		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}
		fmt.Printf("response: %s\n", resp.Message)
	}
}
