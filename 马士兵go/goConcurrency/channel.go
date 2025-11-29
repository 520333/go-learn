package goConcurrency

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}
func ChannelOperate() {
	var ch = make(chan int) //无缓冲channel
	go func() {
		ch <- add(10, 20)
	}()
	go func() {
		v := <-ch
		fmt.Println("Received from channel, value is:", v)
	}()
	time.Sleep(time.Second)
}
