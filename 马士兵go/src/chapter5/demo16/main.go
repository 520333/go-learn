package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T", now)
	fmt.Println(now)
	fmt.Printf("年：%v\n", now.Year())
	fmt.Printf("月：%v\n", int(now.Month()))
	fmt.Printf("日：%v\n", now.Day())
	fmt.Println("----------------")
	fmt.Printf("当前年月日:%d-%d-%d 时分秒:%d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())

	datestr := fmt.Sprintf("当前年月日:%d-%d-%d 时分秒:%d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("20060102"))
}
