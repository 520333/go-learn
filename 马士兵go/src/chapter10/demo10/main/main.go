package main

import (
	"chapter10/demo10/model"
	"fmt"
)

func main() {
	// s := model.Student{"丽丽", 21}
	// var s *model.Student = &model.Student{"丽丽", 21}
	// fmt.Println(s)
	s := model.NewStudent("丽丽", 10)
	fmt.Println(*s)
}
