package main

import "fmt"

func main() {
	var a []int = []int{1, 4, 7, 3, 6, 9}
	var b []int = make([]int, 10)
	copy(b, a)
	fmt.Println(b)
	b[1] = 99
	fmt.Println(b)
	fmt.Println(a)
}
