package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("nub=%d\r\n", num)
	}
}

// 单向使用channel
func main() {
	// var ch1 chan int       //双向channel
	// var ch2 chan<- float64 //单向channel
	// var ch3 <-chan int     //单向channel只能读取
	// c := make(chan int, 3)
	// var send chan<- int = c //send-only
	// var read <-chan int = c //recv-only
	// send <- 1
	// <-read
	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(time.Second * 10)
}
