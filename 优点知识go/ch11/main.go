package main

import "fmt"

var g int = 100

func sum(a, b int) int {
	var c = 10 // c局部变量
	return a + b + c + g
}

func funcValRef(a int) {
	a = 1000
	fmt.Printf("in func inner:%d\n", a)
}

func funcValRefPoint(a *int) {
	*a = 1000
	fmt.Printf("in funcValRefPoint inner:%d\n", *a)
}
func main() {
	var a, b, c int
	a = 10
	b = 20
	c = a + b + g
	fmt.Println(a, b, c)
	fmt.Println(sum(a, b))

	var p int = 100
	funcValRef(p)
	fmt.Printf("in func main:%d\n", p)

	funcValRefPoint(&p)
	fmt.Printf("in func main:%d\n", p)

}
