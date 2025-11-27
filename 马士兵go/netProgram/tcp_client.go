package netProgram

import (
	"encoding/gob"
	"encoding/json"
	"log"
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
