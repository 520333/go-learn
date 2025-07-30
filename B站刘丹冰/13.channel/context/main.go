package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var stop bool // 共享变量

func cpuInfo(ctx context.Context) {
	fmt.Printf("traceid:%s\r\n", ctx.Value("traceid")) // 拿到一个请求的id
	// 记录一些日志，
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出CPU监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu信息")
		}
	}

}
func main() {

	wg.Add(1)
	// context提供了三种函数WithCancel、WithTimeout、WithValue

	// 1.手动
	// ctx1, cancel := context.WithCancel(context.Background())
	// ctx2, _ := context.WithCancel(ctx1)

	// 2.主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	// 3.时间点cancel
	// ctx2, _ := context.WithDeadline(context.Background(), time.Now())

	// 4.WithValue
	valueCtx := context.WithValue(ctx, "traceid", "qjw12j")
	go cpuInfo(valueCtx)
	// time.Sleep(time.Second * 6)
	// cancel()
	wg.Wait()
	fmt.Println("监控完成")
}
