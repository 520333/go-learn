package main

import "fmt"

// 闭包函数
func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}
func main() {
	// 闭包特性
	next := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(next())
	}
	fmt.Println("==============")
	next1 := autoIncrement()
	for i := 0; i < 3; i++ {
		fmt.Println(next1())
	}
}
