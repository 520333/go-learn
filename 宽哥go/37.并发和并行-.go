package main

// import (
// 	"fmt"
// 	"time"
// )

// // 做包子函数
// myfunc makeBuns(filling string) {
// 	startTime := time.Now()
// 	fmt.Printf("%s馅,开始的时间:\n", filling, startTime)

// 	fmt.Printf("开始做%s馅的包子。\n", filling)
// 	fmt.Printf("开始剁%s馅...\n", filling)
// 	fmt.Println("开始擀皮...")
// 	time.Sleep(time.Second)

// 	fmt.Printf("开始包%s馅的包子...\n", filling)
// 	fmt.Printf("开始蒸%s馅的包子...\n", filling)
// 	fmt.Println("======================")
// 	cost := time.Since(startTime) //计算花费时间
// 	fmt.Printf("%s馅,共耗费时间:\n", filling, cost)

// }

// myfunc main() {
// 	/*
// 		协程: 2KB    线程: 8MB
// 		协程不需要关系底层逻辑，内存管理及垃圾回收
// 	*/
// 	fillings := []string{"韭菜", "鸡蛋", "西葫芦"}
// 	for _, v := range fillings { //_忽略索引
// 		go makeBuns(v)
// 	}
// 	time.Sleep(time.Second * 3) //开发时不推荐
// }
