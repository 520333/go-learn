package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex
var wg sync.WaitGroup

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("数据读取中...")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("数据写入中...")
	time.Sleep(time.Second * 10)
	fmt.Println("写入数据成功")
	lock.Unlock()
}
func main() {
	wg.Add(6)
	go write()
	for i := 0; i < 5; i++ {
		go read()
	}
	wg.Wait()
}
