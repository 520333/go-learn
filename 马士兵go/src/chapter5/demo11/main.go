package main

import "fmt"

var Func01 = func(num1, num2 int) int {
	return num1 * num2
}

func main() {
	var num1, num2 int = 20, 20
	reuslt := func(num1, num2 int) int {
		return num1 + num2
	}(num1, num2)
	fmt.Println(reuslt)
	sub := func(num1, num2 int) int {
		return num1 - num2
	}
	fmt.Println(sub(5, 15))
	fmt.Println(Func01(5, 2))
}
