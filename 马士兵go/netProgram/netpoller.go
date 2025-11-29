package netProgram

import (
	"log"
	"net"
	"sync"
	"time"
)

// BIONet 网络IO阻塞(使用系统调用syscall的IO)
func BIONet() {
	addr := "127.0.0.1:5678"
	var wg = sync.WaitGroup{}
	// 1模拟读 体会读的阻塞状态
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		buf := make([]byte, 1024)
		log.Println("start read:", time.Now().Format("15:04:05.000"))
		n, _ := conn.Read(buf)
		log.Println("content:", string(buf[:n]), time.Now().Format("15:04:05.000"))
	}(&wg)

	wg.Add(1)
	// 2模拟写
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		l, _ := net.Listen("tcp", addr)
		defer l.Close()
		for {
			conn, _ := l.Accept()
			go func(conn net.Conn) {
				defer conn.Close()
				log.Println("connected")
				time.Sleep(time.Second * 3)
				conn.Write([]byte("Bloking I/O"))
			}(conn)
		}
	}(&wg)
	wg.Wait()
}

// BIOChannel channel go自管理数据的IO阻塞
func BIOChannel() {
	var wg = sync.WaitGroup{}
	var ch = make(chan struct{})
	// 1模拟读
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("start read:", time.Now().Format("15:04:05.000"))
		content := <-ch
		log.Println("content:", content, time.Now().Format("15:04:05.000"))
	}(&wg)

	wg.Add(1)
	// 2模拟写
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}(&wg)
	wg.Wait()
}
