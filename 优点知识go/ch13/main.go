package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] //2,3,4,5
	fmt.Println(s1)
	s2 := arr[:6]
	fmt.Println(s2)
	s3 := arr[2:]
	fmt.Println(s3)
	s4 := arr[:]
	fmt.Println(s4)
	fmt.Println(len(arr), cap(arr))

	s5 := make([]int, 5, 10)
	fmt.Printf("s5=%v len=%d cap=%d\n", s5, len(s5), cap(s5))

	s6 := arr[2:6]
	fmt.Printf("s6=%v len=%d cap=%d\n", s6, len(s6), cap(s6))
}
