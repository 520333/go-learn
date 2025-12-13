package main

import (
	"fmt"
	"time"
)

func screw(c chan int) {
	i := 1
	for {
		fmt.Printf("正在拧第\033[1;32;40m%d\033[0m个螺丝\n", i)
		c <- i
		i++
		time.Sleep(time.Second)
	}
}

// myfunc main() {
// 	//定义拧螺丝通道
// 	screwChan := make(chan int, 100)
// 	defer close(screwChan)
// 	//定义一个关闭通道的通道
// 	stop := make(chan bool)
// 	defer close(stop)
// 	go screw(screwChan)
// 	go myfunc() {
// 		time.Sleep(10 * time.Second)
// 		fmt.Println("下班了，停止拧螺丝")
// 		stop <- true
// 	}()
// 	for {
// 		select {
// 		case <-stop:
// 			//说明下班时间到，并且往stop通道内发送true数据
// 			return
// 		case s := <-screwChan:
// 			fmt.Printf("第\033[1;32;40m%d\033[0m个螺丝完成。\n", s)
// 		}
// 	}
// }
