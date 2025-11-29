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
	defer udpConn.Close()
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
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send %s(%v) from to %s\n", string(data), wn, raddr)
	log.Println(udpConn.RemoteAddr())

}
