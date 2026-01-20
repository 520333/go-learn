package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloWorld struct{}

func (hw *HelloWorld) HelloWorld(req string, resp *string) error {
	*resp = req + "你好!"
	return nil
}
func main() {
	// 1.注册RPC服务，绑定对象方法
	if err := rpc.RegisterName("hello", &HelloWorld{}); err != nil {
		fmt.Println("注册rpc服务失败:", err)
		return
	}
	// 2.设置监听
	listener, err := net.Listen("tcp", ":8004")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("listening port on :8004")
	// 3.建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		defer conn.Close()
		fmt.Println("new connection")
		// 4.绑定服务
		go rpc.ServeConn(conn)
	}
}
