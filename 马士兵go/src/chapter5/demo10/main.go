package main

import (
	"chapter5/demo10/testutils"
	"fmt"
)

var num int = test()

func test() int {
	fmt.Println("test function execution.")
	return 10
}

func init() {
	fmt.Println("main init function execution...")
}
func main() {
	fmt.Println("main function execution")
	fmt.Printf("Age=%v Sex=%v Name=%v\n", testutils.Age, testutils.Sex, testutils.Name)
}
