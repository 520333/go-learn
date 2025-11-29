package goConcurrency

import (
	"fmt"
	"os"
	"os/signal"
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

func SelectChannelSignal() {
	go func() {
		for true {
			fmt.Println(time.Now().Format("15:04:05.000"))
			time.Sleep(time.Millisecond * 300)
		}
	}()
	//select {}
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal)
	//signal.Notify(chSignal, os.Interrupt)
	//signal.Notify(chSignal, os.Interrupt, os.Kill)
	select {
	case <-chSignal:
		fmt.Println("receive os signal: Interrupt")
	}

}
