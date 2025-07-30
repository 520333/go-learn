package main

import (
	"fmt"
	"sync"
)

// WaitGroup主要用于goroutine的执行，等到Add方法要和Done方法配套结束后
func main() {
	var wg sync.WaitGroup
	wg.Add(100) // 监控goroutine

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()

		}(i)
	}
	wg.Wait() // 等待100个Add结束否则阻塞

}
