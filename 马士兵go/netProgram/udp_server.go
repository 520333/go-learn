package netProgram

import (
	"log"
	"net"
	"os"
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

// UdpReceiverBroadCast 广播接收端
func UdpReceiverBroadCast() {
	// 1广播监听地址
	laddr, _ := net.ResolveUDPAddr(udp, ":6789")

	// 2广播监听
	udpConn, _ := net.ListenUDP(udp, laddr)

	// 3接受数据
	// 4处理数据
	buf := make([]byte, 1024)
	for {
		rn, raddr, _ := udpConn.ReadFromUDP(buf)
		log.Printf("receives from %s data %s\n", string(buf[:rn]), raddr)
	}
}

// UdpFileUploadServer UDP文件上传
func UdpFileUploadServer() {
	// 1.建立UDP连接
	lAddress := ":5678"
	lAddr, _ := net.ResolveUDPAddr(udp, lAddress)
	udpConn, _ := net.ListenUDP(udp, lAddr)
	defer func(udpConn *net.UDPConn) {
		_ = udpConn.Close()
	}(udpConn)
	log.Printf("%s server is listening on %s\n", udp, udpConn.LocalAddr())

	// 2.接受文件名并确认
	buf := make([]byte, 4*1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fileName := string(buf[:rn])
	if _, err := udpConn.WriteToUDP([]byte("fileName ok"), raddr); err != nil {
		log.Fatalln(err)
	}
	// 3.接受文件内容，并写入文件
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	i := 0
	for {
		rn, _, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		// 写入文件
		if _, err := file.Write(buf[:rn]); err != nil {
			log.Println(err)
		}
		i++
		log.Println("file write some content", i)
	}
}
