package main

import (
	"context"
	"fmt"
	"gateway/base/proxy/grpc_proxy/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var msg = "this is client"

func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:8005",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.NewClient err: %v", err)
	}
	defer conn.Close()
	c := proto.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("------------ UnaryEcho Client ------------")
	reply, err := c.UnaryEcho(ctx, &proto.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("client.UnaryEcho err: %v", err)
	} else {
		fmt.Println(reply)
	}
}
