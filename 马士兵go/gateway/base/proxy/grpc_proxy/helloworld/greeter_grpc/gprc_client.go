package main

import (
	"context"
	"fmt"
	"gateway/base/proxy/grpc_proxy/helloworld/greeter_grpc/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1.连接gRPC服务
	//grpcConn, err := grpc.Dial("127.0.0.1:8080")
	grpcConn, err := grpc.NewClient(
		"127.0.0.1:8004",
		// 抑制安全策略
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("grpc dial error:", err)
		return
	}
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			fmt.Println("grpc conn close error:", err)
		}
	}(grpcConn)

	// 2.初始化客户端
	grpcClient := pb.NewHelloServiceClient(grpcConn)

	// 3.调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := grpcClient.Hello(ctx, &pb.Person{Name: "海绵宝宝", Age: 25})
	if err != nil {
		fmt.Println("grpc client call error:", err)
		return
	}
	fmt.Println(reply)
}
