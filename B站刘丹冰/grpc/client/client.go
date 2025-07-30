package main

import (
	"context"
	"fmt"
	"learn/grpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// grpc.WithInsecure() 已经弃用
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "dawn",
		G:    proto.Gender_FEMALE,
		Mp: map[string]string{
			"name":    "dawn",
			"company": "宝哥云计算",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Messgae)
}
