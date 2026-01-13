package udp

import (
	"fmt"
	"net"
)

func Server() {
	// 1.监听服务器指定端口
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080})
	if err != nil {
		fmt.Printf("listen udp error: %v\n", err)
	}
	fmt.Println("udp server is started!")

	// 2.读取客户端数据
	var data [1024]byte
	n, clientAddr, err := conn.ReadFromUDP(data[:])
	if err != nil {
		fmt.Printf("read error,clientAddr: %v err:%v\n", clientAddr, err)
	}
	fmt.Printf("read from clientAddr: %v data: %v count: %v\n", clientAddr, string(data[:n]), n)

	// 3.处理请求并响应
	_, err = conn.WriteToUDP([]byte("received success!"), clientAddr)
	if err != nil {
		fmt.Printf("write to udp error: %v\n", err)
	}
}
