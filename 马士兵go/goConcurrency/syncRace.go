package goConcurrency

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	var counter = 0
	var gs = 100
	wg.Add(gs)
	lock := &sync.Mutex{}
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 1000; k++ {
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
