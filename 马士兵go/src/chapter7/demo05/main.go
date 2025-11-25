package main

import "fmt"

func main() {
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [3]int{3, 2, 1}
	fmt.Println(arr2)

	var arr3 = [...]int{4, 5, 6}
	fmt.Println(arr3)

	var arr4 = [...]int{2: 66, 0: 33, 1: 99, 3: 88}
	fmt.Println(arr4)
}
