package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁-解决资源竞争 做到原子化
*/
// var total int
var total int32
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, 1)
		// lock.Lock()
		// total += 1 //竞争
		// lock.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, -1)
		// lock.Lock()
		// total -= 1 //竞争
		// lock.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total) //没有锁结果变得随机
}
