package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "golang你好"
	fmt.Println(len(str))
	for i, v := range str {
		fmt.Printf("索引：%v 值：%c\n", i, v)
	}
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	count := strings.Count("golangandjavaga", "ga")
	fmt.Println(count)
	flag := strings.EqualFold("hello", "Hello")
	fmt.Println(flag)

	fmt.Println("hello" == "HELLO")

	flag1 := strings.Index("golangandjavaga", "ga")
	fmt.Println(flag1)
}
