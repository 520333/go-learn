package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) Set(name string, age int) {
	fmt.Println("调用Set()方法")
	s.Name = name
	s.Age = age
}
func TestStudent(i interface{}) {
	val := reflect.ValueOf(i)
	fmt.Println(val)

	n := val.Elem().NumField()
	fmt.Printf("结构体字段数:%v\n", n)
	val.Elem().Field(0).SetString("海绵宝宝")
	val.Elem().Field(1).SetInt(199)
}
func main() {
	var s Student = Student{"派大星", 18}
	TestStudent(&s)
	fmt.Println("最终 s =", s)
}
