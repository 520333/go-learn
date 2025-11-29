package goConcurrency

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func GoroutineGo() {
	// 定义输出奇数的函数
	printOdd := func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	// 定义输出偶数的函数
	printEven := func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	// 顺序调用
	//printOdd()
	//printEven()
	// 并发调用
	go printOdd()

	go printEven()
	time.Sleep(time.Second)

}

func GoroutineWG() {
	wg := sync.WaitGroup{}
	// 定义输出奇数的函数
	printOdd := func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	// 定义输出偶数的函数
	printEven := func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
	wg.Add(2)
	go printOdd()
	go printEven()
	go func() {
		wg.Wait()
		fmt.Println("wait in inner goroutine")
	}()
	wg.Wait()
	fmt.Println("after main wait")
}

func GoroutineRandom() {
	wg := sync.WaitGroup{}
	workersNum := 10
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go func(n int) {
			wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func GoroutineNum() {
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()
	//启动大量goroutine
	for {
		go func() {
			fmt.Println("in goroutineNum")
			time.Sleep(100 * time.Second)
		}()
	}

}

// GoroutineAnts 使用ants三方包控制goroutine
func GoroutineAnts() {
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()
	size := 1024
	pool, err := ants.NewPool(size)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Release()
	for {
		err := pool.Submit(func() {
			v := make([]byte, 1024)
			_ = v
			//fmt.Println("in goroutine")
			time.Sleep(100 * time.Millisecond)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
