package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	for {
		time.Sleep(time.Second)
		fmt.Println("dawn")
	}
}
func main() {
	go asyncPrint() // 主程序退出 协程也退出 主死随从

	for i := 0; i < 100; i++ {
		// 匿名函数启动goroutine
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
