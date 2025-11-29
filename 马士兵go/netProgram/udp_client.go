package netProgram

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

// UdpFileUploadClient 文件传输(上传)
func UdpFileUploadClient() {
	// 1.获取文件信息
	fileName := "D:/D/zhaochuan.mp3"
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	// 获取文件
	fileInfo, _ := file.Stat()
	//fileInfo.Size(), fileInfo.Name()
	log.Println("send file size:", fileInfo.Size())
	// 2.连接服务器
	rAddress := "192.168.50.100:5678"
	rAddr, _ := net.ResolveUDPAddr(udp, rAddress)
	udpConn, _ := net.DialUDP(udp, nil, rAddr)
	defer func(udpConn *net.UDPConn) {
		_ = udpConn.Close()
	}(udpConn)

	// 3.发送文件名字
	if _, err := udpConn.Write([]byte(fileInfo.Name())); err != nil {
		log.Fatalln(err)
	}

	// 4.服务端确认
	buf := make([]byte, 4*1024)
	rn, _ := udpConn.Read(buf)
	if "fileName ok" != string(buf[:rn]) {
		log.Fatalln(errors.New("server not ready"))
	}
	// 5.发送文件内容
	i := 0
	for {
		rn, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		if _, err := udpConn.Write(buf[:rn]); err != nil {
			log.Fatalln(err)
		}
		i++
	}
	log.Println(i)
	log.Println("file upload complete.")
	time.Sleep(time.Second * 2)
}
