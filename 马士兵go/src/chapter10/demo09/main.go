package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func method01(s Student) {
	fmt.Println(s.Name, reflect.TypeOf(s.Name))
}

func method02(s *Student) {
	fmt.Println((*s).Name, reflect.TypeOf(s.Name))
}

func (s Student) test01() {
	fmt.Println(s.Name)
}
func (s *Student) test02() {
	fmt.Println((*s).Name)
}
func main() {
	var s Student = Student{"珊珊"}
	method01(s)
	method02(&s)

	s.test01()
	(&s).test02()
	s.test02()
}
