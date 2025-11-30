package goConcurrency

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
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

func SyncLockAndNo() {
	wg := &sync.WaitGroup{}
	var counter = 0
	var gs = 1000
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
	wg.Add(1)
	//var lck2 sync.Mutex
	go func() {
		defer wg.Done()
		for k := 0; k < 10000; k++ {
			//lck2.Lock()
			counter++
			//lck2.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func SyncMutex() {
	wg := &sync.WaitGroup{}
	var loc sync.Mutex
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Before lock:", n)
			loc.Lock()
			fmt.Println("locked:", n)
			time.Sleep(time.Second)
			loc.Unlock()
			fmt.Println("After lock:", n)
		}(i)
	}
	wg.Wait()
}

func SyncRLock() {
	wg := &sync.WaitGroup{}
	var rwLock sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwLock.RLock()
			fmt.Println(time.Now())
			time.Sleep(time.Second)
			rwLock.RUnlock()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwLock.Lock()
			fmt.Println(time.Now(), "Lock")
			time.Sleep(time.Second)
			rwLock.Unlock()
		}()
	}
	wg.Wait()
}

func SyncMapErr() {
	m := map[string]int{}
	go func() {
		for {
			m["key"] = 0
		}
	}()
	go func() {
		for {
			_ = m["key"]
		}
	}()
	select {}
}
func SyncMap() {
	var m sync.Map
	go func() {
		for {
			m.Store("key", 1)
		}
	}()
	go func() {
		for {
			_, _ = m.Load("key")
		}
	}()
	select {}
}
func SyncMapMethod() {
	var wg sync.WaitGroup
	var m sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(n, fmt.Sprintf("value:(%d)", n))
		}(i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(m.Load(n))
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
	m.Delete(4)
}

func SyncAtomicAdd() {
	counter := atomic.Int32{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.Load())
}
func SyncAtomicValue() {
	var loadConfig = func() map[string]string {
		return map[string]string{
			"title":   "Go云原生",
			"varConf": fmt.Sprintf("%d", rand.Int31()),
		}
	}
	var config atomic.Value
	go func() {
		for {
			config.Store(loadConfig())
			fmt.Println("latest config was loaded", time.Now().Format("15:04:05.99999999"))
			time.Sleep(time.Second)
		}
	}()
	for {
		go func() {
			c := config.Load()
			fmt.Println(c, time.Now().Format("15:04:05.99999999"))
		}()
		time.Sleep(time.Millisecond * 400)
	}
	select {}
}

func SyncPool() {
	// 原子的计数器
	var counter int32 = 0
	//定义元素的Newer,创建器
	elementNewer := func() any {
		atomic.AddInt32(&counter, 1)
		return new(bytes.Buffer)
	}
	// Pool的初始化
	var pool = sync.Pool{
		New: elementNewer,
	}
	// 并发的申请和交回元素
	workerNum := 1024
	wg := &sync.WaitGroup{}
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			defer wg.Done()
			buffer := pool.Get().(*bytes.Buffer)
			//buffer := elementNewer().(*bytes.Buffer)
			defer pool.Put(buffer)
			_ = buffer.String()
		}()
	}
	wg.Wait()
	fmt.Println("elementNewer:", counter)
}

func SyncOnce() {
	config := make(map[string]string)
	once := sync.Once{}
	loadConfig := func() {
		//once.Do(func() {
		config = map[string]string{
			"varInt": fmt.Sprintf("%d", rand.Int31()),
		}
		fmt.Println("config loaded")
		//})
	}
	workers := 10
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			once.Do(func() {
				loadConfig()
			})

			_ = config
		}()
	}
	wg.Wait()
}

func SyncCond() {
	wg := &sync.WaitGroup{}
	var data []int
	dataLen := 1024 * 1024
	cond := sync.NewCond(&sync.Mutex{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < dataLen; i++ {
			data = append(data, i*i)
		}
		cond.Broadcast()
		fmt.Println("cond broadcast, len(data)=", len(data))
	}()

	const workers = 8
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			cond.L.Lock()
			for len(data) < dataLen {
				cond.Wait()
			}
			fmt.Println("处理数据,数据长度:", len(data))
			cond.L.Unlock()
		}()
	}
	wg.Wait()
}
