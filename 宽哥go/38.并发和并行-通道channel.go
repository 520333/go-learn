package main

// myfunc makeBuns(filling string, buns chan string) { // 做包子函数
// 	fmt.Printf("开始做%s馅的包子。\n", filling)
// 	fmt.Printf("开始剁%s馅...\n", filling)
// 	fmt.Println("开始擀皮...")
// 	time.Sleep(time.Second * 1)

// 	fmt.Printf("开始包%s馅的包子...\n", filling)
// 	fmt.Printf("开始蒸%s馅的包子...\n", filling)

// 	time.Sleep(time.Second * 1)
// 	fmt.Printf("%s馅的包子已经蒸好了，可以上菜了 时间时间:%s\n", filling, time.Now())

// 	//通道也是一种类型
// 	//发送数据: channel1 <- "数据"
// 	buns <- filling
// 	//取数据: data1 := <- channel1

// }
// myfunc waiter(buns chan string) { //上菜函数
// 	bun := <-buns
// 	fmt.Printf("上菜:%s馅包子  上菜时间:%s\n", bun, time.Now())
// }

// myfunc main() {
// 	/*
// 		协程: 2KB    线程: 8MB
// 		协程不需要关系底层逻辑，内存管理及垃圾回收
// 	*/
// 	buns := make(chan string, 5)
// 	defer close(buns) //关闭通道
// 	fillings := []string{"韭菜", "鸡蛋", "西葫芦"}
// 	for _, v := range fillings { //_忽略索引
// 		go makeBuns(v, buns)
// 	}

// 	for i := 0; i < len(fillings); i++ {
// 		// go waiter(buns)
// 		time.Sleep(time.Second * 3)
// 		bun := <-buns
// 		fmt.Printf("上菜:%s馅包子  上菜时间:%s\n", bun, time.Now())
// 	}
// 	// time.Sleep(time.Second * 3) //开发时不推荐
// }
