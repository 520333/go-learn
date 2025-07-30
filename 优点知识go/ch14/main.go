package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 10, 20, 30, 40, 50)
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))
	fmt.Println(arr, len(arr), cap(arr))

	var s2 []int
	fmt.Println(s2, s2 == nil, len(s2), cap(s2))
	s2 = append(s2, 0)
	fmt.Println(s2, s2 == nil, len(s2), cap(s2))
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		fmt.Println(s2, s2 == nil, len(s2), cap(s2))
	}

	var s3 = make([]int, 1024)
	fmt.Println(s3, len(s3), cap(s3))

	s3 = append(s3, 200)
	fmt.Println(s3, len(s3), cap(s3)) //容量=1024+1024/2 =1536

	s4 := []int{1, 2, 3, 4}
	s5 := make([]int, 10)
	fmt.Println(s4, s5)
	copy(s5, s4)
	fmt.Println(s4, s5)

	s6 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("before delete slice s6=%v\n", s6)
	s6 = append(s6[:4], s6[5:]...) //删除第4个元素
	fmt.Printf("after delete slice s6=%v\n", s6)
}
