package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func read(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("读取:%v\n", v)
		time.Sleep(time.Millisecond * 100)
	}
}
func write(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("写入数据:%v\n", i)
	}
	close(ch)
}

func main() {
	var intChan = make(chan int, 50)
	wg.Add(2)
	go write(intChan)
	// go read(intChan)
	wg.Wait()
}
