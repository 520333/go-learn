package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	name := "海绵宝宝"
	nameP := &name
	fmt.Println("变量的值为：", name)
	fmt.Println("变量的内存地址为：", nameP, "指针大小：", len(*nameP))

	nameV := *nameP
	fmt.Println("nameP指针指向的值为：", nameV)

	*nameP = "公众号：海绵宝宝抓水母"
	fmt.Println("指针指向变量的值：", *nameP)
	fmt.Println("name变量的值：", name)
	var intP *int = new(int)
	*intP = 10
	// 简略写法
	intP1 := new(int)
	*intP1, _ = strconv.Atoi("1")
	fmt.Println(*intP1)

	age := 18
	modifyAge(&age)
	fmt.Println(age)
	var w io.Writer = os.Stdout
	wp := &w
	fmt.Println(wp)
}

func modifyAge(age *int) {
	*age = 20
}
