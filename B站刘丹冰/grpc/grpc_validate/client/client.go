package main

import (
	"context"
	"fmt"
	"learn/grpc/grpc_validate/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// client拦截器
func Interceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	md := metadata.New(map[string]string{
		"appid":  "1010101",
		"appkey": "chuangdevops",
	})

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Since 计算耗时
	time.Sleep(time.Second)
	fmt.Printf("耗时：%s\r\n", time.Since(start))
	return err

}

type customCredential struct{}

func main() {
	var opts []grpc.DialOption
	// opts = append(opts, grpc.WithUnaryInterceptor(Interceptor))
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "dawn@qq",
		Mobile: "14411112222",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Email)
}
