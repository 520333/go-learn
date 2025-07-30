package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, name string) {
	for {
		select {
		// case <-stopCh:
		case <-ctx.Done():
			fmt.Println(name, "停止指令已经收到,马上停止")
			return
		default:
			fmt.Println(name, "正在监控......")
		}
		time.Sleep(1 * time.Second)
	}

}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户ID为：", userId)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	// stopCh := make(chan bool)
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		// watchDog(stopCh, "[监控狗1]")
		watchDog(ctx, "[监控狗1]")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗2]")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗3]")
	}()

	valCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second)
	stop()
	// stopCh <- true
	wg.Wait()

}
