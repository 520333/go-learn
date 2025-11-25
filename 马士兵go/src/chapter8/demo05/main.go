package main

import "fmt"

func main() {
	var intarr [6]int = [6]int{1, 4, 7, 3, 6, 9}
	var slice []int = intarr[1:4]
	fmt.Println(len(slice))
	fmt.Println(slice)
	slice2 := append(slice, 12)
	// slice2 = append(slice, 13)
	// slice2 = append(slice, 14)
	fmt.Println(intarr)
	fmt.Println(slice)
	fmt.Println(slice2)
	slice3 := []int{99, 44}
	slice = append(slice, slice3...)
	fmt.Println(slice)
}
