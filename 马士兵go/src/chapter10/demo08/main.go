package main

import "fmt"

type Student struct {
	Name string
}

func (s *Student) test1() {
	fmt.Println(s.Name)
}
func method01(s Student) {
	fmt.Println(s.Name)
}
func main() {
	s := Student{"海绵宝宝"}
	fmt.Println(&s)
	method01(Student{"bbb"})
}
