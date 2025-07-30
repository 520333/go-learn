package main

import (
	"fmt"
	"learn/grpc_stream/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type Server struct {
}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v %d", time.Now().Unix(), i),
		})
		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
