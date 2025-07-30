package main

import (
	"fmt"
	"learn/ch26-gotesting/calculate"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d is even ? %v \n", i, calculate.Even(i))

	}
}
