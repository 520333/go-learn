package netProgram

import (
	"fmt"
	"log"
	"net"
	"time"
)

func UdpClientBasic() {
	//建立连接
	raddr, err := net.ResolveUDPAddr(udp, ":9876")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP(udp, nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	//写
	data := []byte("Go udp Program")
	wn, err := udpConn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) to %s\n", string(data), wn, raddr)

	//读
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
}
func UdpClientConnect() {
	//建立连接
	raddr, err := net.ResolveUDPAddr(udp, ":9876")
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP(udp, nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(udpConn.RemoteAddr())

	//写
	data := []byte("Go udp Program")
	wn, err := udpConn.Write(data)
	//wn, err := udpConn.WriteToUDP(data, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) to %s\n", string(data), wn, raddr)

	//读
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
}

// 对等连接
func UdpClientPeer() {
	//解析地址
	laddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:6789")
	if err != nil {
		log.Fatalln(err)
	}
	//监听
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s server is listening on %s\n", udp, udpConn.LocalAddr())
	defer func(udpConn *net.UDPConn) {
		_ = udpConn.Close()
	}(udpConn)
	//远程地址
	raddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:9876")
	if err != nil {
		log.Fatalln(err)
	}
	//写
	data := []byte("Go udp Program")
	wn, err := udpConn.WriteToUDP(data, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) to %s\n", string(data), wn, raddr)

	//读
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
}

// UdpSenderMultiCast 多播发送端
func UdpSenderMultiCast() {
	// 1建立连接
	address := "224.1.1.2:6789"
	raddr, err := net.ResolveUDPAddr(udp, address)
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP(udp, nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	// 2发送数据 循环发送
	for {
		data := fmt.Sprintf("[%s]: %s", time.Now().Format("15:04:05.000"), "hello")
		wn, err := udpConn.Write([]byte(data))
		if err != nil {
			log.Println(err)
		}
		log.Printf("send %s(%v) to %s\n", string(data), wn, raddr)
		time.Sleep(time.Second)
	}

}

// UdpSenderBroadCast 广播发送端
func UdpSenderBroadCast() {
	// 1监听地址
	// 2建立连接
	laddr, _ := net.ResolveUDPAddr(udp, ":9876")
	udpConn, _ := net.ListenUDP(udp, laddr)
	// 3发送数据
	rAddress := "192.168.50.255:6789"
	raddr, _ := net.ResolveUDPAddr(udp, rAddress)
	for {
		data := fmt.Sprintf("[%s]: %s", time.Now().Format("15:04:05.000"), "hello!!!")
		wn, _ := udpConn.WriteToUDP([]byte(data), raddr)
		log.Printf("send %s(%v) to %s\n", string(data), wn, raddr)
		time.Sleep(time.Second)
	}
}
