package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = (10+20)%3 + 3 - 7
	fmt.Println(num1, num2)

	var a int = 8
	var b int = 4
	fmt.Printf("a=%v,b=%v \n", a, b)
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("a=%v,b=%v \n", a, b)

	fmt.Println(5 == 9)
	fmt.Println(5 != 9)
	fmt.Println(5 > 9)
	fmt.Println(5 < 9)

}
