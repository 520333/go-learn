package main

/*
channel适用场景:
	1.消息传递，消息过滤
	2.信号广播
	3.事件订阅和广播
	4.任务分发
	5.结果汇总
	6.并发控制
	7.同步和异步
*/
/*
import (
	"fmt"
	"time"
)

// 不要通过共享内存来通信，而要通过通信来实现内存共享
func main() {
	// 无缓冲channel适用于通知、B要第一时间知道是否已经完成
	var msg chan string
	msg = make(chan string, 0) // chan的初始化值如果是0放值进去会被阻塞 无缓冲channel
	go func(msg chan string) {
		data := <-msg // 取值
		fmt.Println(data)
	}(msg)
	msg <- "dawn" // 放值
	time.Sleep(time.Second * 2)

	// 有缓冲channel适用于消费者和生产者之间的通信
	msg2 := make(chan string, 1) // 容量大于1 有缓冲channel
	msg2 <- "chuang"
	fmt.Println(<-msg2)
}
*/
