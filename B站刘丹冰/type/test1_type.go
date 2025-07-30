package main

import (
	"fmt"
	"strconv"
)

type Myint2 int

func (mi Myint2) string() string {
	return strconv.Itoa(int(mi))
}
func main() {
	type Myint = int // 定义int类型的别名
	var i Myint
	var j int = 8
	fmt.Println(i + j)
	fmt.Printf("%T\r\n", i)

	type Meint int // 自定义类型
	var x Meint = 12
	fmt.Println(x + Meint(j))

	var y Myint2 = 12
	fmt.Println(y.string())
	fmt.Printf("%T\r\n", y)

	var a interface{} = "abcd"
	switch a.(type) {
	case string:
		fmt.Println("字符串")
	}
	m := a.(string)
	fmt.Printf("%T", m)
}
