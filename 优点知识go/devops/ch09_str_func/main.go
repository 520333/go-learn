package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str = "this is str"
	fmt.Println(len(str))

	var str1 = "你好"
	var str2 = "golang"
	fmt.Println(str1 + ":" + str2)
	var str3 = fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)

	var str4 = "123.456.789"
	var arr = strings.Split(str4, ".")
	fmt.Println(arr)

	var str5 = strings.Join(arr, "----")
	fmt.Println(str5)

	var str6 = "hello 张三"
	for i := 0; i < len(str6); i++ {
		fmt.Println(string(str6[i]))
	}

	for _, value := range str6 {
		fmt.Println(string(value))
	}
	var num int = 20
	var str7 = strconv.Itoa(num)
	fmt.Printf("%T %v", str7, str7)

	var num2 float64 = 20.11213312
	var str8 = strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Println(str8)
	var str9 = "100"
	var num3, _ = strconv.Atoi(str9)
	fmt.Printf("%T %v", num3, num3)
}
