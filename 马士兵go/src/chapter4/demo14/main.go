package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d\n", i)
		if i == 14 {
			return
		}
	}
}
