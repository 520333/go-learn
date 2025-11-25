package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		intChan <- 10
	}()
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "hello golang"
	}()
	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-stringChan:
		fmt.Println(v)
	default:
		fmt.Println("防止select被阻塞")
	}
}
