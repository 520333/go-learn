package main

import "fmt"

func main() {
	strs := "我是谁abc"
	a := []rune(strs)
	b := []byte(strs)
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %d\n", b, b)

	//单引号
	var a1 = 'a'
	fmt.Printf("%T %v", a1, a1)
}
