package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享的资源
var (
	sum   int
	mutex sync.Mutex
)
var rwmutex sync.RWMutex

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i

}

func readSum() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	b := sum
	return b
}

func run() {
	// 开启100个协程让sum+=10
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	wg.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only Once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑...")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	wg.Wait()
}

func syncMap() {
	syncMap := sync.Map{}
	syncMap.Store(1, 1)
	syncMap.Store(1, 2)
	fmt.Println(syncMap.Load(1))
}

func main() {
	run()
	doOnce()
	race()
	syncMap()

}
