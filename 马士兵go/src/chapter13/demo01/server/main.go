package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("读取数据失败:%v\n", err)
			return
		}
		fmt.Printf("客户端消息:%v\n", string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务端启动...")
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("监听失败:%v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("等待连接失败:%v", err)
		} else {
			fmt.Printf("连接成功,conn=%v 客户端信息:%v\n", conn, conn.RemoteAddr())
		}
		go process(conn)
	}
}
