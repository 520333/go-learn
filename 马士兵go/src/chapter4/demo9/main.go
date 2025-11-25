package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 300 {
			break
		}
		fmt.Printf("%d\n", sum)
	}

}
