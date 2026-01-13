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
	// 1.与服务器建立TCP连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer conn.Close()
	// 2.接受服务端响应
	go getServerData(conn)
	//clientBuf := make([]byte, 1024)
	//n, _ := conn.Read(clientBuf)
	//fmt.Println("from client:", string(clientBuf[:n]))

	// 3.向服务端发送消息
	//conn.Write([]byte("hello world"))
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

func getServerData(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		data := strings.TrimSpace(string(buf[:n]))
		if data != "" {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "服务端:", string(buf[:n]))
		}
	}
}
