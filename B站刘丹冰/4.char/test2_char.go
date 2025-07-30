package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串比较
	a := "hello"
	b := "bello"
	fmt.Println(a != b) // 是否相等
	fmt.Println(a > b)  // 比较大小
	var bb byte = 'b'
	var h byte = 'h'
	fmt.Println(bb, h)

	// 字符串是否包含
	name := "宝哥云原生必修课：golang go"
	AA := strings.Contains(name, strings.ToLower("GO"))
	fmt.Println(AA)

	// 字符串出现次数
	fmt.Println(strings.Count(name, "go"))

	// 字符串分割
	fmt.Println(strings.Split(name, "："))

	// 字符串是否包含前缀
	fmt.Println(strings.HasPrefix(name, "宝"))

	// 字符串是否包含后缀
	fmt.Println(strings.HasSuffix(name, "go"))

	// 字符串index出现的位置
	fmt.Println(strings.IndexRune(name, []rune(name)[8]))

	// 字符串替换 0不替换  1替换一个  2替换2个  -1全部替换
	fmt.Println(strings.Replace(name, "go", "java", -1))

	// 大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))

	// 去掉左右两侧指定字符
	fmt.Println(strings.Trim("#?hello 1 1#$", "#$?"))

}
