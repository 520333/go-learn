package main

import "fmt"

func main() {
	var arr = [3]int16{1, 2, 3}
	fmt.Println(len(arr))
	fmt.Printf("arr的地址为：%p\n", &arr)
	fmt.Printf("arr[0]的地址为：%p\n", &arr[0])
	fmt.Printf("arr[1]的地址为：%p\n", &arr[1])
	fmt.Printf("arr[2]的地址为：%p\n", &arr[2])
}
