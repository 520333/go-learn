package main

import "fmt"

func main() {

	fmt.Println("this is a myfunc")
	defer func() {
		if r := recover(); r != nil { // recover用于捕获panic
			fmt.Println("recover if A", r)
		}
	}()
	panic("this ist a panic")   // panic 会导致程序退出
	var names map[string]string // 未初始化也会造成程序退出
	names["go"] = "go工程师"

}
