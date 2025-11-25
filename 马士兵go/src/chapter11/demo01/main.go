package main

import (
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("test.txt")
	if err != nil {
		fmt.Printf("文件打开出错，对应错误为%v \n", err)
	}
	fmt.Println(file)
	if err := file.Close(); err != nil {
		fmt.Println("关闭失败")
	}
}
