package netProgram

import (
	"log"
	"net"
)

const udp = "udp"

func UdpServerBasic() {
	//解析地址
	laddr, err := net.ResolveUDPAddr(udp, ":9876")
	if err != nil {
		log.Fatalln(err)
	}
	//监听
	udpConn, err := net.ListenUDP(udp, laddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s server is listening on %s\n", udp, udpConn.LocalAddr())
	defer udpConn.Close()
	//读操作
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
	//写操作
	data := []byte("received:" + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) from to %s\n", string(data), wn, raddr)
}

func UdpServerConnect() {
	//解析地址
	laddr, err := net.ResolveUDPAddr(udp, ":9876")
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
	log.Println(udpConn.RemoteAddr())
	//读操作
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
	//写操作
	data := []byte("received:" + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, raddr)
	//wn, err := udpConn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) from to %s\n", string(data), wn, raddr)
	log.Println(udpConn.RemoteAddr())

}

// 对等连接
func UdpServerPeer() {
	//解析地址
	laddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:9876")
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
	raddr, err := net.ResolveUDPAddr(udp, "127.0.0.1:6789")
	if err != nil {
		log.Fatalln(err)
	}
	//读操作
	buf := make([]byte, 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
	//写操作
	data := []byte("received:" + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) from to %s\n", string(data), wn, raddr)
}

// UdpReceiverMultiCast 多播接收端
func UdpReceiverMultiCast() {
	// 1组播监听地址
	address := "224.1.1.2:6789"
	gaddr, err := net.ResolveUDPAddr(udp, address)
	if err != nil {
		log.Fatalln(err)
	}

	// 2组播监听
	udpConn, err := net.ListenMulticastUDP(udp, nil, gaddr)
	if err != nil {
		log.Fatalln(err)
	}
	// 3接收数据 循环接受
	buf := make([]byte, 1024)
	for {
		rn, raddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		log.Printf("receives %s from data %s\n", string(buf[:rn]), raddr)
	}
}
