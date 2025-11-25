package main

import "fmt"

func main() {
	//var name string
	//fmt.Println("请输入您的名字:")
	//fmt.Scanln(&name)
	//
	//var age int
	//fmt.Println("请输入您的年龄:")
	//fmt.Scanln(&age)
	//
	//var score float32
	//fmt.Println("请输入您的成绩:")
	//fmt.Scanln(&score)
	//
	//var isVIP bool
	//fmt.Println("是否VIP用户:")
	//fmt.Scanln(&isVIP)
	//
	//fmt.Println(name, age, score, isVIP)
	fmt.Println("请输入员工姓名、年龄、成绩、是否VIP。使用空格进行分割")
	var name string
	var age int
	var score float32
	var isVIP bool
	_, _ = fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
	fmt.Println(name, age, score, isVIP)

}
