package main

import (
	"fmt"
	"sync"
)

type safeInt struct {
	a    int
	lock sync.Mutex
}

func (si *safeInt) increment() {
	si.lock.Lock()

	defer si.lock.Unlock()
	si.a++
}
func (si *safeInt) Get() int {
	si.lock.Lock()
	defer si.lock.Unlock()
	return si.a

}

func main() {
	var si safeInt
	// var a int
	// a++
	// go func() {
	// 	a++
	// }()
	// time.Sleep(time.Millisecond)
	// fmt.Println(a)
	si.increment()
	var wg sync.WaitGroup

	for i := 0; i < 60000020; i++ {
		wg.Add(1)
		go func() {
			si.increment()
			wg.Done()
		}()
		wg.Wait()
	}

	// time.Sleep(time.Millisecond)
	fmt.Println(si.Get())

}
