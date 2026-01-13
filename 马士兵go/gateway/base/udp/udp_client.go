package udp

import (
	"fmt"
	"net"
)

func Client() {
	// 1.建立udp连接
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("connection failed. err:%v\n", err)
	}
	// 2.发送数据给服务器
	data := "Hello udp server!"
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Printf("connection failed. err:%v\n", err)
	}

	//3.接受服务器数据
	var result = make([]byte, 1024)
	n, ra, err := conn.ReadFromUDP(result)
	if err != nil {
		fmt.Println("receive failed. err:", err)
	}
	fmt.Printf("response from server, addr:%v data:%v\n", ra, string(result[:n]))

}
