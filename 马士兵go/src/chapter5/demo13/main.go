package main

import "fmt"

func add(num1, num2 int) int {
	defer fmt.Printf("num1=%v\n", num1)
	defer fmt.Printf("num2=%v\n", num2)
	num1 += 90
	num2 += 50
	var sum int = num1 + num2
	fmt.Printf("sum=%v\n", sum)

	return sum
}
func main() {
	add(30, 60)
}
