package goConcurrency

import (
	"fmt"
	"sync"
	"time"
)

func SelectChannelCloseSignal() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		fmt.Println("发出信号,close(ch)")
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		for {
			select {
			case <-ch:
				fmt.Println("收到停止信号，<-ch")
				return
			default:
			}

			fmt.Println("业务逻辑处理中。。。", i)
			i++
			time.Sleep(time.Millisecond * 300)
		}
	}()
	wg.Wait()
}
