package main

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误已经捕获:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
func main() {
	test()
	fmt.Println("test函数执行成功。。。")
	fmt.Println("正常执行下面的逻辑")
}
