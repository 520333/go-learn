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
