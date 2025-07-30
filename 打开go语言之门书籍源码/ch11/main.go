package main

import (
	"fmt"
	"sync"
	"time"
)

// 工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

// 工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装（" + c + "）"
		}
	}()
	return out
}

// 工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包【" + c + "】"
		}
	}()
	return out
}

// 封装一个组件加快生产
func megre(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	// 扇入，需要启动多个goroutine用于处理多个channel中的数据
	for _, cs := range ins {
		go p(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 洗菜
func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

// 烧水
func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧开的水"
	}()
	return water
}
func main() {
	result := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
	coms := buy(100)
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)

	phones := megre(phones1, phones2, phones3)
	packs := pack(phones)
	for p := range packs {
		fmt.Println(p)
	}
	vegetablesCh := washVegetables()
	waterCh := boilWater()
	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了,可以做火锅了:", vegetables, water)
}
