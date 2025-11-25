package model

import "fmt"

type person struct {
	Name string
	age  int
}

// 定义工厂模式的函数
func NewPerson(name string) *person {
	return &person{Name: name}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("对不起传入的年龄范围不正确!")
	}
}
func (p *person) GetAge() int {
	return p.age
}
