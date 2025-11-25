package main

import "fmt"

func main() {
	var intarr [6]int = [6]int{1, 4, 7, 2, 5, 8}
	var slice []int = intarr[1:4]
	fmt.Println(slice, len(slice), cap(slice))

	slice2 := slice[1:2]
	fmt.Println(slice2)
}
