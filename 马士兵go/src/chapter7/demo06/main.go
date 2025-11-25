package main

import "fmt"

func main() {
	// var arr1 = [3]int{3, 6, 9}
	// fmt.Println(arr1)
	// fmt.Printf("类型%T\n", arr1)
	// var arr2 = [6]int{3, 6, 9, 1, 4, 7}
	// fmt.Println(arr2)
	// fmt.Printf("类型%T\n", arr2)
	var arr3 = [3]int{3, 6, 9}
	test1(&arr3)
	fmt.Println(arr3)
}
func test1(arr *[3]int) {
	(*arr)[0] = 7
}
