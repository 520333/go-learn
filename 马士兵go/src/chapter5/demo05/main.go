package main

import "fmt"

func test(nums ...int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
func main() {
	test(1000)
	fmt.Println("=================")
	test(1, 2, 3, 4, 5)
}
