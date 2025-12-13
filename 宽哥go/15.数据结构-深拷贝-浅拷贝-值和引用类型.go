package main

// myfunc main() {
// 	s1 := "dawn"
// 	s2 := s1
// 	fmt.Println(s1, s2) // stdout: dawn dawn
// 	s2 = "chuang"
// 	fmt.Println(s1, s2) // stdout: dawn chuang

// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := slice1
// 	fmt.Println("slice1:", slice1) // stdout: slice1: [1 2 3 4 5]
// 	fmt.Println("slice2:", slice2) // stdout: slice2: [1 2 3 4 5]

// 	// slice2[1] = 88
// 	// fmt.Println("slice1:", slice1) // stdout: slice1: [1 88 3 4 5] 浅拷贝 引用同一个内存地址 所以slice1的值也被更改
// 	// fmt.Println("slice2:", slice2) // stdout: slice2: [1 88 3 4 5]

// 	slice3 := make([]int, len(slice1), cap(slice1))
// 	copy(slice3, slice1)
// 	fmt.Println("slice3:", slice3)
// 	slice3[1] = 88
// 	fmt.Println("slice3:", slice3)                       // stdout: slice3: [1 88 3 4 5]
// 	fmt.Println("slice1:", slice1)                       // stdout: slice1: [1 2 3 4 5]
// 	fmt.Println("slice1的内存地址:", unsafe.Pointer(&slice1)) //stdout: slice1的内存地址: 0xc000008078
// 	fmt.Println("slice3的内存地址:", unsafe.Pointer(&slice3)) //stdout: slice3的内存地址: 0xc0000080d8

// 	fmt.Println("slice1的第一个元素的内存地址:", unsafe.Pointer(&slice1[1])) //stdout: slice1的第一个元素的内存地址: 0xc00000e458
// 	fmt.Println("slice2的第一个元素的内存地址:", unsafe.Pointer(&slice2[1])) //stdout: slice1的第一个元素的内存地址: 0xc00000e458
// 	fmt.Println("slice3的第一个元素的内存地址:", unsafe.Pointer(&slice3[1])) //stdout: slice3的第一个元素的内存地址: 0xc00000e4e8

// }
