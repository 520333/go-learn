package main

/*
import (
	"fmt"
)

myfunc main() {
	c := make(chan int)
	go myfunc() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //close可以关闭一个channel
	}()

	for {
		// 如果ok为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok { // if是判断最后一个ok是否为true
			fmt.Println(data)
		} else {
			break
		}
	}
	// 简写
	for data := range c {
		fmt.Println(data)
	}

	// timeNow := time.Now()
	// fmt.Printf("当前时间:%s\n", timeNow.Format("2006-01-02 15:04:05"))
	fmt.Println("Main Finished...")
}
*/
