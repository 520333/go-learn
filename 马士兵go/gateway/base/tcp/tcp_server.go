package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// 1.监听服务器端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen failed. err:", err)
		return
	}
	fmt.Println("服务器监听已启动 :8080")
	// 2.创建TCP连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept failed. err:", err)
	}
	conn.Write([]byte("received success!"))
	fmt.Println("客户端已连接...")
	defer conn.Close()

	// 3.处理客户端请求，打印到控制台
	go GetClientData(conn)
	// 4.对客户端进行响应
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("write failed, err:", err)
			break
		}
	}
}

func GetClientData(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		data := strings.TrimSpace(string(buf[:n]))
		if data != "" {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "客户端:", string(buf[:n]))
		}
	}
}
