package main

import (
	"context"
	"fmt"
	"learn/grpc_stream/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "宝哥云计算"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a, a.Data)
	}

	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("宝哥%d", i),
		})
		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "我是客户端"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
