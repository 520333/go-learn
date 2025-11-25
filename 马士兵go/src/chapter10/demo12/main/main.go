package main

import (
	"chapter10/demo12/model"
	"fmt"
)

func main() {
	//创建person结构体的示例：
	p := model.NewPerson("丽丽")
	p.SetAge(200)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
}
