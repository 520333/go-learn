package main

import (
	"fmt"
)

func main() {
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	fmt.Println(cap(intarr))
	slice := intarr[1:3]
	fmt.Println("slice", slice)
	fmt.Println("slice的元素个数:", len(slice))
	fmt.Println("slice的容量个数:", cap(slice))
	fmt.Printf("slice内存地址:%p\n", &slice)
	fmt.Printf("slice[0]内存地址:%p\n", &slice[0])
	fmt.Printf("slice[1]内存地址:%p\n", &slice[1])

	var arr = [3]int{1, 2, 3}
	slice2 := arr[1:2]
	fmt.Println(slice2)
	fmt.Printf("arr内存地址:%p\n", &arr)
	fmt.Printf("arr[1]内存地址:%p\n", &arr[1])

	fmt.Printf("slice内存地址:%p\n", &slice2)
	fmt.Printf("slice[1]内存地址:%p\n", &slice2[0])
	slice2[0] = 20
	fmt.Println(arr)

}
