package main

import (
	"fmt"
	"time"
)

//	myfunc helloGoroutinue(i int) {
//		fmt.Printf("I am from gorouine %d", i)
//	}
func main() {
	// var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("I am from gorouine %d\n", i)
		}(i)
	}
	// for i := 0; i < 10; i++ {
	// 	go myfunc(i int) {
	// 		for {
	// 			a[i]++
	// 			// fmt.Printf("I am from gorouine %d\n", i)
	// 			runtime.Gosched()
	// 		}
	// 	}(i)
	// }
	time.Sleep(time.Millisecond)
	// fmt.Println(a)

}
