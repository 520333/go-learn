package main

import (
	"learn/rpc/new_hello_rpc/handler"
	"learn/rpc/new_hello_rpc/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	//1.实例化server
	listener, _ := net.Listen("tcp", ":1234")

	//2.注册处理逻辑handler
	server_proxy.RegisterHelloService(&handler.NewHelloService{})

	//3.启动rpc服务
	for {
		conn, _ := listener.Accept() //建立socket套接字
		go rpc.ServeConn(conn)
	}

}
