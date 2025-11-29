package goConcurrency

import (
	"fmt"
	"runtime"
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

func ChannelGoroutineNumCtr() {
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 500)
		}
	}()
	const size = 1024
	ch := make(chan struct{}, size)
	for {
		ch <- struct{}{}
		go func() {
			time.Sleep(time.Second * 10)
			<-ch
		}()
	}
}
