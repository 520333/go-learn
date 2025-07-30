package main

import "fmt"

func changeEle(arr *[3]int) {
	arr[0] = 100
}
func main() {
	var arr1 [5]int
	arr1[1] = 10
	arr2 := []int{1, 2, 3}
	fmt.Println(arr1, arr2, len(arr2), cap(arr2))

	arr3 := [...]int{4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr3, len(arr3))

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	for i := range arr3 {
		fmt.Println(i)
	}

	arr4 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i, v := range arr4 {
		fmt.Println(i, v)
	}
	arr5 := [3]int{1, 2, 3}
	changeEle(&arr5)

	fmt.Println(arr5)

}
