package main

// myfunc main() {
// 	// 定义：var 切片名称 []切片类型
// 	var s1 []int
// 	fmt.Println("数据:", s1, "默认长度:", len(s1), "默认容量:", cap(s1)) //sdtout: [] 默认长度: 0 默认容量: 0
// 	s1 = append(s1, 111, 222)
// 	// fmt.Println("")
// 	fmt.Println("数据:", s1, "默认长度:", len(s1), "默认容量:", cap(s1)) //sdtout: [] 默认长度: 0 默认容量: 0

// 	//指定长度
// 	s2 := make([]int, 5, 10)
// 	fmt.Println("切片长度:", len(s2), "切片容量:", cap(s2)) //stdout: 切片长度: 5 切片容量: 10
// 	fmt.Println("数据:", s2)                          //stdout: 数据: [0 0 0 0 0]

// 	s2 = append(s2, 1, 2, 3, 4, 5, 6)
// 	fmt.Println("数据:", s2) //stdout: 数据: [0 0 0 0 0]

// 	// 通过append增加长度超过10 容量自动*2
// 	fmt.Println("切片长度:", len(s2), "切片容量:", cap(s2)) //stdout: 切片长度: 5 切片容量: 20

// 	s2[0] = 99
// 	fmt.Println("数据:", s2) //stdout: 数据: [99 0 0 0 0 1 2 3 4 5 6]

// 	for k, v := range s2 {
// 		fmt.Printf("第%d个数据是:%d\n", k+1, v)
// 	}

// }
