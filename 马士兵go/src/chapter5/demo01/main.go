package main

import "fmt"

func cal(num1 int, num2 int) (int, int) {
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	return num1 + num2, num1 - num2
}
func main() {
	cal(-5, 2)
}
