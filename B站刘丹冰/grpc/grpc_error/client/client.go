package main

import (
	"context"
	"fmt"
	"learn/grpc/grpc_error/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3) //设置超时
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "dawn"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message(), st.Code())
	}
}
