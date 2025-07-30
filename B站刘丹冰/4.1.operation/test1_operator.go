package main

import (
	"fmt"
)

func main() {
	// 运算符: + - * / % ++ --
	var a, b = 1, 2
	fmt.Println(a + b)
	var astr, bstr = "hello", "dawn"
	fmt.Println(astr + bstr)
	fmt.Println(4 % 2) //取余
	a++
	fmt.Println("a++ = ", a)
	a = a - 1 // a--, a = a-1
	fmt.Println("a-- = ", a)

	// 逻辑运算符: && || !
	var abool, bbool = true, false
	if abool && bbool { // 逻辑与: abool和bbool两个必须都为真
		fmt.Println("逻辑与ok")
	}
	if abool || bbool { // 逻辑或: abool和bbool有一个为真
		fmt.Println("逻辑或ok")
	}
	if !bbool { // 逻辑非: abool取反
		fmt.Println("逻辑非ok")
	}

	// 位运算符: & | ^ << >> 常用于高性能需求开发 原理：将数值转换成二进制后做运算
	var a1, a2 = 60, 13
	// 按位与
	// 60 = 0011 1100
	// 13 = 0000 1101
	fmt.Println(a1 & a2) // 12 = 0000 1100

	// 按位或
	// 60 = 0011 1100
	// 13 = 0000 1101
	fmt.Println(a1 | a2) // 61 = 0011 1101

	// 按位异或
	// 60 = 0011 1100
	// 13 = 0000 1101
	fmt.Println(a1 ^ a2)  // 49 = 0011 0001
	fmt.Println(a1 << a2) // 49 = 0011 0001
	fmt.Println(a1 >> a2) // 49 = 0011 0001
}
