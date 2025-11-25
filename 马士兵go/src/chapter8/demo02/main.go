package main

import "fmt"

func main() {
	slice := make([]int, 4, 20)
	fmt.Println(slice, cap(slice))
	slice[0] = 66
	slice[1] = 88
	fmt.Println(slice)
	slice = append(slice, 12)
	fmt.Println(slice)
	for k, v := range slice {
		fmt.Println(k, v)
	}
}
