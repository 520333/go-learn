package main

import (
	"context"
	"fmt"
	"learn/grpc/interpretor/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// client拦截器
func Interceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Since 计算耗时
	time.Sleep(time.Second)
	fmt.Printf("耗时：%s\r\n", time.Since(start))
	return err

}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(Interceptor))
	// opt := grpc.WithUnaryInterceptor(Interceptor)
	// grpc.WithInsecure() 已经弃用
	// conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "dawn"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Messgae)
}
