package main

import "fmt"

func main() {
	var arr = [2][3]int16{{0, 2, 3}, {1, 21, 2}}
	fmt.Printf("arr[0]的地址:%p\n", &arr[0])
	fmt.Printf("arr[0][0]的地址:%p\n", &arr[0][0])
	fmt.Printf("arr[0][1]的地址:%p\n", &arr[0][1])
	fmt.Printf("arr[0][2]的地址:%p\n", &arr[0][2])
	fmt.Println("----------------------------")
	fmt.Printf("arr[1]的地址:%p\n", &arr[1])
	fmt.Printf("arr[1][0]的地址:%p\n", &arr[1][0])
	fmt.Printf("arr[1][1]的地址:%p\n", &arr[1][1])
	fmt.Printf("arr[1][2]的地址:%p\n", &arr[1][2])

	var arr1 [2][4]int = [2][4]int{{1, 4, 7}, {2, 5, 8}}
	fmt.Println(arr1)
}
