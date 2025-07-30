package main

import "fmt"

// go语言提供了数组、切片(slice)、map、list
func main() {
	// var courses1 [3]string // 只有3个元素的数组
	// var courses2 [4]string // 只有4个元素的数组
	// //[]string 和 [3]string 是两种不同的类型
	// courses1[0] = "go"
	// courses1[1] = "grpc"
	// courses1[2] = "gin"
	// fmt.Printf("%T\r\n", courses1)
	// fmt.Printf("%T\r\n", courses2)

	// for _, value := range courses1 {
	// 	fmt.Println(value)
	// }

	// 数组初始化
	//var courses1 [3]string = [3]string{"go", "grpc", "gin"}
	courses1 := [3]string{"go", "grpc", "gin"} // 可以简化为

	for _, value := range courses1 {
		fmt.Println(value)
	}
	courses2 := [3]string{2: "gin"}
	for _, value := range courses2 {
		fmt.Println(value)
	}
	courses3 := [...]string{"go", "grpc", "gin"}
	for _, value := range courses3 {
		fmt.Println(value)
	}
	for i := 0; i < len(courses3); i++ {
		fmt.Println(courses3[i])
	}
	courses4 := [...]string{"go", "grpc", "gin"}
	if courses3 == courses4 {
		fmt.Println("equal")
	}

	// 多维数组
	var coursesInfo [3][4]string
	coursesInfo[0] = [4]string{"go", "1h", "dawn", "快速入门"}
	// coursesInfo[0][0] = "go"
	// coursesInfo[0][1] = "java"
	// coursesInfo[0][2] = "php"
	coursesInfo[1] = [4]string{"grpc", "2h", "dawn", "grpc快速入门"}
	coursesInfo[2] = [4]string{"gin", "1.5h", "dawn", "gin高级开发"}
	fmt.Println(len(coursesInfo))
	for i := 0; i < len(coursesInfo); i++ {
		for j := 0; j < len(coursesInfo[i]); j++ {
			fmt.Print(coursesInfo[i][j] + " ")
		}
		fmt.Println()
	}
	for _, row := range coursesInfo {
		for _, column := range row {
			fmt.Print(column + " ")
		}
		fmt.Println()
		fmt.Println(row)
	}

}
