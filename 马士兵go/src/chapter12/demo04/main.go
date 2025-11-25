package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int64
var wg sync.WaitGroup
var mu sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		// mu.Lock()
		// total += 1
		// mu.Unlock()
		atomic.AddInt64(&total, 1)
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		// mu.Lock()
		// total -= 1
		// mu.Unlock()
		atomic.AddInt64(&total, -1)

	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
