package goConcurrency

import (
	"fmt"
	"sync"
)

func SyncErr() {
	var counter = 0
	var gs = 100
	wg := &sync.WaitGroup{}
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
func SyncLock() {
	var counter = 0
	var gs = 1000
	wg := &sync.WaitGroup{}
	wg.Add(gs)
	lock := &sync.Mutex{}
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
