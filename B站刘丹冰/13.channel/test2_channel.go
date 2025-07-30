package main

/*
import (
	"fmt"
	"time"
)

// 有缓冲chanel
func main() {
	c := make(chan int, 3) // 这是一个有缓冲通道可以容纳最多3个元素
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行:len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		num := <-c //从c中接收并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束...")
}
*/
