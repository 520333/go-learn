package main

import (
	"fmt"
	"time"
)

// var done bool
// var lock sync.Mutex
var done = make(chan struct{})

func g1(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	// lock.Lock()
	// defer lock.Unlock()
	// done = true
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	// lock.Lock()
	// defer lock.Unlock()
	// done = true
	ch <- struct{}{}
}
func main() {
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	go g1(g1Channel)
	go g2(g2Channel)
	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			time.Sleep(10 * time.Millisecond)
			fmt.Println("timtout")
			return
		}
	}
	// for {
	// 	if done {
	// 		fmt.Println("done")
	// 		time.Sleep(10 * time.Millisecond)
	// 		return
	// 	}
	// }
}
