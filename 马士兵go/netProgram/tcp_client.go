package netProgram

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net"
	"sync"
	"time"
)

const tcp = "tcp"

// 基本读写操作
func TcpClient() {
	//2.模拟多客户端
	//3.并发的客户端请求
	//server_address := "127.0.0.1:5678"
	server_address := ":5678"
	num := 10
	var wg = sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			//1.建立连接
			conn, err := net.Dial(tcp, server_address)
			if err != nil {
				log.Fatalln(err)
				return
			}
			defer conn.Close()
			log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
		}(&wg)
	}
	wg.Wait()
}
func TcpTimeoutClient() {
	server_address := "192.168.1.100:5678"
	num := 10
	var wg = sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.DialTimeout(tcp, server_address, time.Second)
			if err != nil {
				log.Fatalln(err)
				return
			}
			defer conn.Close()
			log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
		}(&wg)
	}
	wg.Wait()
}

func TcpBacklogClient() {
	server_address := ":5678"
	num := 256
	var wg = sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(wg *sync.WaitGroup, n int) {
			defer wg.Done()
			conn, err := net.DialTimeout(tcp, server_address, time.Second)
			if err != nil {
				log.Fatalln(err)
				return
			}
			defer conn.Close()
			log.Printf("%d dial connection establish,client addr %s\n", n, conn.LocalAddr())
		}(&wg, i)
		time.Sleep(time.Millisecond * 30)
	}
	wg.Wait()
}
func TcpClientRW() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	// 从服务端接受数据,Read
	//buf := make([]byte, 1024)
	//rn, err := conn.Read(buf)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("received from server data is :", string(buf[:rn]))

	// 向服务端发送数据,Write
	wn, err := conn.Write([]byte("send some data from client" + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("server write len is %d bytes\n", wn)
}
func TcpWClient() {
	server_address := ":5678"
	conn, err := net.Dial(tcp, server_address)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	conn.LocalAddr()
	//select {}
	// 阻塞操作
	//buf := make([]byte, 1024)
	//rn, err := conn.Read(buf)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("received from server data is :", string(buf[:rn]))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	for {
		buf := make([]byte, 10)
		rn, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("received from server data is :", string(buf[:rn]))
	}

}

// 并发
func TcpClientRWConcurrency() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	conn.LocalAddr()

	wg := sync.WaitGroup{}

	//并发的写
	wg.Add(1)
	go CliWrite(conn, &wg)
	//并发的读
	wg.Add(1)
	go CliRead(conn, &wg)
	wg.Wait()

}
func CliWrite(conn net.Conn, wg *sync.WaitGroup) {
	//向客户端发送数据Write
	defer wg.Done()
	for {
		wn, err := conn.Write([]byte("send some data from client" + "\n"))
		if err != nil {
			log.Println(err)
		}
		log.Printf("client write len is %d \n", wn)
		time.Sleep(time.Millisecond * 3000)
	}
}
func CliRead(conn net.Conn, wg *sync.WaitGroup) {
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

// 格式化输出
func TcpClientFormat() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	conn.LocalAddr()

	wg := sync.WaitGroup{}

	//接收端
	wg.Add(1)
	go CliReadFormat(conn, &wg)
	wg.Wait()

}
func CliReadFormat(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// 从客户端接受数据Read
		// 从服务端接收数据
		type Message struct {
			Id      uint   `json:"id,omitempty"`
			Code    string `json:"code,omitempty"`
			Content string `json:"content,omitempty"`
		}
		var message = Message{}
		//1.JSON 文本解码
		decoder := json.NewDecoder(conn)
		if err := decoder.Decode(&message); err != nil {
			log.Println(err)
			continue
		}
		log.Println("json", message)

		//2.GOB 二进制解码
		g := gob.NewDecoder(conn)
		if err := g.Decode(&message); err != nil {
			log.Println(err)
			continue
		}
		log.Println("gob", message)

	}
}

// 短连接
func TcpClientShort() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	conn.LocalAddr()

	wg := sync.WaitGroup{}

	//接收端
	wg.Add(1)
	go CliReadShort(conn, &wg)
	wg.Wait()

}
func CliReadShort(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	type Message struct {
		Id      uint   `json:"id,omitempty"`
		Code    string `json:"code,omitempty"`
		Content string `json:"content,omitempty"`
	}
	var message = Message{}
	for {
		// 从客户端接受数据Read
		// 从服务端接收数据
		//2.GOB 二进制解码
		g := gob.NewDecoder(conn)
		err := g.Decode(&message)
		if err != nil && errors.Is(err, io.EOF) {
			log.Println(err)
			log.Println("link was closed")
			break
		}
		log.Println("gob", message)
	}
}

// 响应服务端的心跳检测
type MessageHB struct {
	ID      uint      `json:"id,omitempty"`
	Code    string    `json:"code,omitempty"`
	Content string    `json:"content,omitempty"`
	Time    time.Time `json:"time,omitempty"`
}

func TcpClientHB() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())
	conn.LocalAddr()
	wg := sync.WaitGroup{}

	//接收端
	wg.Add(1)
	go CliReadPing(conn, &wg)
	wg.Wait()

}
func CliReadPing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	var message = MessageHB{}
	for {
		g := gob.NewDecoder(conn)
		err := g.Decode(&message)
		if err != nil && errors.Is(err, io.EOF) {
			log.Println(err)
			break
		}
		//判断是否为ping类型消息
		if message.Code == "PING-SERVER" {
			log.Println("receive ping from", conn.RemoteAddr())
			CliWritePong(conn, message)
		}
	}
}
func CliWritePong(conn net.Conn, pingMsg MessageHB) {

	var pongMsg = MessageHB{
		ID:      uint(rand.Int()),
		Code:    "PONG-CLIENT",
		Content: fmt.Sprintf("pingID:%v", pingMsg.ID),
		Time:    time.Now(),
	}

	g := gob.NewEncoder(conn)
	if err := g.Encode(pongMsg); err != nil {
		log.Println(err)
		return
	}
	log.Println("pong was send to", conn.RemoteAddr())
	return
}

// 连接池客户端
func TcpClientPool() {
	server_address := ":5678"
	//建立连接池
	pool, err := NewTcpPool(server_address, PoolConfig{
		Factory:     &TcpConnFactory{},
		InitConnNum: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("pool created", pool.Len())
	wg := sync.WaitGroup{}
	clientNum := 50
	wg.Add(clientNum)
	for i := 0; i < clientNum; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := pool.Get()
			if err != nil {
				log.Println(err)
				return
			}
			//log.Println(conn)
			_ = pool.Put(conn)
		}(&wg)
	}
	wg.Wait()
	//释放连接池
	_ = pool.Release()
	log.Println(pool, pool.idleList)
}

// 粘包现象
func TcpClientSticky() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())

	buf := make([]byte, 1024)
	for {
		rn, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("received data:", string(buf[:rn]))
	}
}

// 粘包现象编解码器
func TcpClientStickyCoder() {
	server_address := ":5678"
	conn, err := net.DialTimeout(tcp, server_address, time.Second)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Printf("dial connection establish,client addr %s\n", conn.LocalAddr())

	decoder := NewDecoder(conn)
	data := ""
	i := 0
	for {
		if err := decoder.Decode(&data); err != nil {
			log.Println(err)
			break
		}
		log.Println(i, "received data:", data)
		i++
	}
}

// TcpClientSpecial TCP专用客户端拨号方法
func TcpClientSpecial() {
	raddr, err := net.ResolveTCPAddr(tcp, ":5678")
	if err != nil {
		log.Fatalln(err)
	}
	tcpConn, err := net.DialTCP(tcp, nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpConn.Close()
	log.Printf("dial connection establish,client addr %s\n", tcpConn.LocalAddr())
	for {
		buf := make([]byte, 1024)
		rn, err := tcpConn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("received data:%v     len:%v", string(buf[:rn]), len(buf[:rn]))
	}
}
