package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("Name=%v Age=%v \n", s.Name, s.Age)
	return str
}

func main() {
	var s Student = Student{
		Name: "海绵宝宝",
		Age:  21,
	}
	fmt.Println(&s)
}
