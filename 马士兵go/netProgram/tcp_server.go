package netProgram

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand/v2"
	"net"
	"sync"
	"time"
)

func TcpServer() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("accept connection from %s\n", conn.RemoteAddr())
	}
}
func TcpBacklogServer() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			log.Printf("accept connection from %s\n", conn.RemoteAddr())
			time.Sleep(time.Second)
		}(conn)
	}
}
func TcpServerRW() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConn(conn)
	}
}
func TcpW() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnW(conn)
	}
}
func HandleConn(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	//向客户端发送数据Write
	wn, err := conn.Write([]byte("send some data from server" + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("server write len is %d bytes\n", wn)

	// 从客户端接受数据Read
	buf := make([]byte, 1024)
	rn, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}
	log.Println("received from server data is :", string(buf[:rn]))

}
func HandleConnW(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	//向客户端发送数据Write
	//data := []byte("send some data from server" + "\n")
	//wn, err := conn.Write(data)
	//if err != nil {
	//	log.Println(err)
	//}
	//// 严谨判断是否写入成功,需要：
	//if err == nil && wn == len(data) {
	//	log.Println("write success")
	//}
	//log.Printf("server write len is %d bytes\n", wn)

	// 2写操作会被阻塞
	//for i := 0; i < 300000; i++ {
	//	data := []byte("send some data from server" + "\n")
	//	wn, err := conn.Write(data)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Printf("%d server write len is %d bytes\n", i, wn)
	//}

	time.Sleep(5 * time.Second)
	data := []byte("send some data from server" + "\n")
	wn, err := conn.Write(data)
	if err == nil && wn == len(data) {
		log.Println("write success")
	}
	log.Printf("server write len is %d bytes\n", wn)
	select {}
}

// 并发的读和写操作，全双工
func TcpServerRWConcurrency() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnConcurrency(conn)
	}
}
func HandleConnConcurrency(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer conn.Close()
	wg := sync.WaitGroup{}

	//并发的写
	wg.Add(1)
	go SerWrite(conn, &wg, "abcd\n")
	wg.Add(1)
	go SerWrite(conn, &wg, "efgh\n")
	wg.Add(1)
	go SerWrite(conn, &wg, "ijkl\n")

	//并发的读
	wg.Add(1)
	go SerRead(conn, &wg)
	wg.Wait()
}
func SerWrite(conn net.Conn, wg *sync.WaitGroup, data string) {
	//向客户端发送数据Write
	defer wg.Done()
	for {
		wn, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(err)
		}
		log.Printf("server write len is %d bytes\n", wn)
		time.Sleep(time.Millisecond * 1000)
	}
}
func SerRead(conn net.Conn, wg *sync.WaitGroup) {
	// 从客户端接受数据Read
	defer wg.Done()
	for {
		buf := make([]byte, 1024)
		rn, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}
		log.Println("received from server data is :", string(buf[:rn]))
	}
}

// 格式化传输
func TcpServerFormat() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnFormat(conn)
	}
}
func HandleConnFormat(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer conn.Close()
	wg := sync.WaitGroup{}

	//发送端写
	wg.Add(1)
	go SerWriteFormat(conn, &wg)

	wg.Wait()
}
func SerWriteFormat(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		//向客户端发送数据Write
		// 数据编码后发送
		type Message struct {
			Id      uint   `json:"id,omitempty"`
			Code    string `json:"code,omitempty"`
			Content string `json:"content,omitempty"`
		}
		var message = Message{Id: uint(rand.Int()), Code: "SERVER-STANDARD", Content: "message from server"}
		//1.JSON 文本编码
		encoder := json.NewEncoder(conn)
		if err := encoder.Encode(message); err != nil {
			log.Println(err)
			continue
		}
		log.Println("message was send json")
		time.Sleep(time.Millisecond * 1000)

		//2.GOB 二进制编码
		g := gob.NewEncoder(conn)
		if err := g.Encode(message); err != nil {
			log.Println(err)
			continue
		}
		log.Println("message was send gob")
		time.Sleep(time.Millisecond * 1000)
	}
}

// 短连接
func TcpServerShort() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnShort(conn)
	}
}
func HandleConnShort(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer conn.Close()
	wg := sync.WaitGroup{}

	//发送端写
	wg.Add(1)
	go SerWriteShort(conn, &wg)
	wg.Wait()
}
func SerWriteShort(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	//向客户端发送数据Write
	// 数据编码后发送
	type Message struct {
		Id      uint   `json:"id,omitempty"`
		Code    string `json:"code,omitempty"`
		Content string `json:"content,omitempty"`
	}
	var message = Message{Id: uint(rand.Int()), Code: "SERVER-STANDARD", Content: "message from server"}

	//2.GOB 二进制编码
	g := gob.NewEncoder(conn)
	if err := g.Encode(message); err != nil {
		log.Println(err)
		return
	}
	log.Println("message was send gob")
	log.Println("link will be closed")
	return
}

// 心跳检测
func TcpServerHB() {
	//基于地址建立监听
	//address := "127.0.0.1:5678"
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listening on %s\n", address)
	//接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnHB(conn)
	}
}
func HandleConnHB(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer func() {
		conn.Close()
		log.Println("connection be closed")
	}()
	//独立的 goroutine 连接建立后周期性发送Ping
	wg := sync.WaitGroup{}
	//发送端写
	wg.Add(1)
	go SerPing(conn, &wg)
	wg.Wait()
}
func SerPing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	//接受Pong
	ctx, cancel := context.WithCancel(context.Background())
	go SerReadPong(conn, ctx)

	const maxPingNum = 3
	var pingErrCounter = 0
	ticker := time.NewTicker(time.Second * 3)
	for t := range ticker.C {
		var pingMsg = MessageHB{ID: uint(rand.Int()), Code: "PING-SERVER", Time: t}
		//2.GOB 二进制编码
		g := gob.NewEncoder(conn)
		if err := g.Encode(pingMsg); err != nil {
			log.Println(err)
			//连接有问题的情况
			pingErrCounter++
			if pingErrCounter == maxPingNum {
				cancel()
				return
			}
		}
		log.Printf("ping send to %v ping id is %v\n", conn.RemoteAddr(), pingMsg.ID)
	}
}
func SerReadPong(conn net.Conn, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var message = MessageHB{}
			g := gob.NewDecoder(conn)
			err := g.Decode(&message)
			if err != nil && errors.Is(err, io.EOF) {
				log.Println(err)
				break
			}
			//判断是否为pong类型消息
			if message.Code == "PONG-CLIENT" {
				log.Printf("receive pong from %s,%v\n", conn.RemoteAddr(), message.Content)
			}
		}
	}
}

// 测试连接池服务端
func TcpServerPool() {
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnPool(conn)
	}
	log.Printf("%v server is listening on %v\n", tcp, listener.Addr())
}
func HandleConnPool(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer func() {
		_ = conn.Close()
		log.Println("connection be closed")
	}()
	select {}
}

// 粘包现象
func TcpServerSticky() {
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("%v server is listening on %v\n", tcp, listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnSticky(conn)
	}
}
func HandleConnSticky(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer func() {
		_ = conn.Close()
		log.Println("connection be closed")
	}()
	//连续发送数据
	data := "package data."
	for i := 0; i < 50; i++ {
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(err)
		}
	}
}

// 粘包现象编解码器
func TcpServerStickyCoder() {
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("%v server is listening on %v\n", tcp, listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go HandleConnStickyCoder(conn)
	}
}
func HandleConnStickyCoder(conn net.Conn) {
	log.Printf("accept connection from %s\n", conn.RemoteAddr())
	defer func() {
		_ = conn.Close()
		log.Println("connection be closed")
	}()
	//连续发送数据
	data := []string{"package data.", "package.", "packa"}
	encoder := NewEncoder(conn)
	for i := 0; i < 50; i++ {
		if err := encoder.Encode(data[rand.IntN(len(data))]); err != nil {
			log.Println(err)
		}
	}
}

// TCP专用方法
func TcpServerSpecial() {
	laddr, err := net.ResolveTCPAddr(tcp, ":5678")
	if err != nil {
		log.Fatalln(err)
	}
	tcpListener, err := net.ListenTCP(tcp, laddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpListener.Close()
	log.Printf("%v server is listening on %v\n", tcp, tcpListener.Addr())
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go HandlerTcpConnSpecial(tcpConn)
	}

}

func HandlerTcpConnSpecial(tcpConn *net.TCPConn) {
	log.Printf("accept connection from %s\n", tcpConn.RemoteAddr())
	data := "tcp message."
	n, err := tcpConn.Write([]byte(data))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("send len:", n)
}
