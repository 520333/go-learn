package main

import (
	"errors"
	"fmt"
)

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}

}
func main() {
	if err := test(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("test函数执行成功。。。")
	fmt.Println("正常执行下面的逻辑")
}
