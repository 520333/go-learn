package main

import (
	"context"
	"fmt"
	"learn/grpc/metadata/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// grpc.WithInsecure() 已经弃用
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	md := metadata.New(map[string]string{
		"name":     "dawn",
		"password": "tingbao89",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "dawn"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Messgae)
}
