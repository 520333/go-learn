package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int
	msg = make(chan int, 2)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all Done...")
	}(msg)
	msg <- 1
	msg <- 2
	close(msg) // 关闭chan
	d := <-msg
	fmt.Println(d)
	// msg <- 3 // chan关闭后无法在放值

	time.Sleep(time.Second * 2)

}
