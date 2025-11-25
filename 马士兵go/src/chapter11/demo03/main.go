package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../demo01/test.txt")
	if err != nil {
		fmt.Printf("打开失败:%v \n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(str) > 0 {
				fmt.Println(str)
			}
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取成功并且全部读取完毕")
}
