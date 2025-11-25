package main

import "fmt"

func main() {
label1:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i=%d  j=%d \n", i, j)
			if i == 2 && j == 2 {
				break label1
			}
		}
	}
}
