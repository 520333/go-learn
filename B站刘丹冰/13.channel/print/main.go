package main

import (
	"fmt"
	"time"
)

/*
使用两个goroutin交替打印序列，一个goroutine打印数字，一个goroutine打印字母最终效果如下：
12AB34CD56EF78GH910
*/
var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}
}
func main() {
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(time.Second * 10)
}
