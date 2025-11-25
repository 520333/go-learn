package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}
func test1(num1 int, num2 float32, f func(i int)) {
	fmt.Println("----test1", num1, num2, f)
}

type myFunc func(int)

func test3(num1 int, num2 float32, f myFunc) {
	fmt.Println("----test1", num1, num2, f)
}
func test4(num1, num2 int) (int, int) {
	var add = num1 + num2
	var sub = num1 - num2
	fmt.Printf("加法：%v 减法：%v\n", add, sub)
	return add, sub
}
func test5(num1, num2 int) (add int, sub int) {
	add = num1 + num2
	sub = num1 - num2
	fmt.Printf("加法：%v 减法：%v\n", add, sub)
	return
}

func main() {
	a := test
	fmt.Printf("a的类型：%T test函数的类型：%T\n", a, test)
	a(123)
	test1(10, 20, a)
	test1(10, 20, test)

	type myInt int
	var num1 myInt
	num1 = 123
	fmt.Printf("num1=%v,type=%T\n", num1, num1)

	var num2 int = 30
	num1 = myInt(num2)
	fmt.Printf("num2=%v,type=%T\n", num2, num2)

	test3(10, 9.8, a)
	test4(40, 20)
	test5(10, 2)
}
