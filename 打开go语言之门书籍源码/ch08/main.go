package main

import (
	"fmt"
	"time"
)

// <- ch 接收 操作符 '<-' 在左边
// ch <- 发送 操作符 '<-' 在右边
func downloadFile(chanName string) string {
	time.Sleep(time.Second)
	return chanName + " :filePath"
}
func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("宝哥无情")
		ch <- "go routine完成了 "
	}()
	fmt.Println("我是main goroutine")
	v := <-ch
	fmt.Println("接受到了chan中的值为:", v)
	cache := make(chan int, 5)
	cache <- 2
	cache <- 3
	close(cache)
	fmt.Println("cache的容量为:", cap(cache), "元素个数为:", len(cache))

	// 单向管道
	// onlySend := make(chan<- int) //只能发送
	// onlyReceive := make(<-chan int) //只能接收

	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)

	}
}
