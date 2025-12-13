package main

import (
	"fmt"
	"time"
)

func main() {
	//用go创建一个形参返回值为空的函数
	/*go myfunc() {
		defer fmt.Println("A defer")

		myfunc() {
			defer fmt.Println("B defer")
			runtime.Goexit() //退出当前goroutine
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()*/
	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
