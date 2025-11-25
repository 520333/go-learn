package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	context, err := ioutil.ReadFile("../demo01/test.txt")
	if err != nil {
		fmt.Printf("读取失败:%v \n", err)
	}
	fmt.Println(string(context))
}
