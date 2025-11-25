package main

import "fmt"

func exeChangeNum(num1, num2 *int) {
	var t int
	t = *num1
	*num1 = *num2
	*num2 = t
}

func main() {
	var num1, num2 int = 10, 20
	fmt.Printf("交换前：num1=%v num2=%v\n", num1, num2)
	exeChangeNum(&num1, &num2)
	fmt.Printf("交换后：num1=%v num2=%v\n", num1, num2)
}
