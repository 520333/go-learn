package main

import "fmt"

func main() {
	a := 20
	if a < 20 {
		fmt.Println("a小于20")
	} else if a == 20 {
		fmt.Println("a等于20")
	} else {
		fmt.Println("a大于等于20")
	}
	fmt.Println("===========================")
	switch a {
	case 20:
		fmt.Println("a=20")
		fallthrough
	case 21:
		fmt.Println("a=21")
	case 22:
		fmt.Println("a=22")
	default:
		fmt.Println("啥都不相等")
	}
}
