package main

import "fmt"

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("hello,world!")
// 	time.Sleep(1 * time.Second)
// }

func add(sum int) int {
	if sum == 1 {
		return 1
	}
	return add(sum-1) + sum
}

func main() {
	fmt.Println(add(10))
}
