package main

import "fmt"

var name string = "chuang" // 定义全局变量 全局变量无法使用 ":=" 语法糖

func printAD() {
	// 定义打印广告函数
	fmt.Println("python点我开始学习")
	fmt.Println("kubernetes云原生技术")
	fmt.Println("Java vue全栈学习")
	fmt.Println("printAD函数调用：", name)

}

// myfunc main() {
// 	fmt.Println("===start===")
// 	printAD()
// 	fmt.Println("==end===")
// 	fmt.Println("main方法调用：", name)
// }
