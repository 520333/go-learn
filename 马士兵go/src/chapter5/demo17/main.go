package main

import (
	"fmt"
)

func main() {
	str := "golang"
	fmt.Println(len(str))

	num := new(int)
	fmt.Printf("num类型：%T num值：%v num内存地址：%v num指针指向的地址：%v", num, num, &num, *num)

}
