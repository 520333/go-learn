package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动...")
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println("创建连接失败:", err)
		return
	}
	fmt.Println("连接成功", conn)
	reader := bufio.NewReader(os.Stdin) //终端标准输入
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("终端输入失败:%v\n", err)
	}

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Printf("连接失败:%v\n", err)
	}
	fmt.Printf("向服务器发送数据成功 一共发送了%d字节数据 \n", n)
}
