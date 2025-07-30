package main

import "fmt"

var b int = 200 // 全局变量作用域优先级低于局部变量 全局变量无法使用:=语法糖
var ok []string // 全局变量定义了可以不被使用
var _ int       // 匿名变量
var (
	name string = "dawn"
	age  int    = 18
	sex  string = "男"
)

// go语言定义了变量必须被使用
func main() {
	var a float64   // 方式1: 只声明变量 不赋值 默认值为0
	var b int = 100 // 方式2: 声明变量并赋值
	var c = 'c'     // 方式3: 省去数据类型 自动匹配变量类型
	d := "ok"       // 方式4: 语法糖
	fmt.Println(a, b, c, d)
	fmt.Println(name, age, sex)

	var user1, user2, user3 = "dawn", 3.14, true
	var test int
	fmt.Println(user1, user2, user3, test)

	// 变量优先级： 类变量 > 局部变量 > 全局变量
	{
		d := "local"
		fmt.Println(d)
	}

}
