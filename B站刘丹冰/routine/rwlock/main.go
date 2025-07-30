package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁本质上是将并行的代码串行化了，使用lock肯定会影响性能
// 即使设计了锁，也应该尽量保证并行
// 2组协程 其中一组负责写数据 另一组负责读数据 web绝大部分都是读多写少
// 虽然有多个goroutine 但是仔细分析 协程之间应该并发执行，读和写应该串行
func main() {
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	// 写的goroutine
	go func() {
		time.Sleep(1 * time.Second)
		defer wg.Done()
		rwlock.Lock() // 加写锁，防止别的写锁和读锁获取
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(5 * time.Second)
	}()

	// 读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() // 加读锁，读锁不会阻止别人的读
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
		}()
	}

	wg.Wait()
}
