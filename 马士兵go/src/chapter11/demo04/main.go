package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("文件打开失败:%v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(fmt.Sprintf("%v  -我是你爹不不不!!!\n", i+1))
	}
	writer.Flush()

	s := os.FileMode(0644).String()
	fmt.Println(s)
}
