package netProgram

import (
	"log"
	"net"
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
