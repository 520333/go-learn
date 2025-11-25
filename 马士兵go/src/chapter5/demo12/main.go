package main

import "fmt"

func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}
