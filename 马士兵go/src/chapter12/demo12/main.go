package main

import (
	"fmt"
	"time"
)

func printNum() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
func devide() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("devide()出现错误", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	go printNum()
	go devide()
	time.Sleep(time.Second * 5)
}
