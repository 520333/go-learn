package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request // 返回值是通过reply的值
	return nil
}
func main() {
	//1.实例化server
	listener, _ := net.Listen("tcp", ":1234")

	//2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	//3.启动rpc服务 for死循环防止server关闭
	for {
		conn, _ := listener.Accept()                    //建立socket套接字
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 使用json编解码  go开启协程并发

	}

}
