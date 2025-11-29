package goConcurrency

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineGo() {
	// 定义输出奇数的函数
	printOdd := func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	// 定义输出偶数的函数
	printEven := func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	// 顺序调用
	//printOdd()
	//printEven()
	// 并发调用
	go printOdd()

	go printEven()
	time.Sleep(time.Second)

}

func GoroutineWG() {
	wg := sync.WaitGroup{}
	// 定义输出奇数的函数

	printOdd := func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	// 定义输出偶数的函数
	printEven := func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	wg.Add(2)
	go printOdd()
	go printEven()
	wg.Wait()
}
