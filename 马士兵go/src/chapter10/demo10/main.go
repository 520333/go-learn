package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var s1 Student = Student{"小李", 22}
	fmt.Println(s1)
	var s2 = Student{Age: 21, Name: "海绵宝宝"}
	fmt.Println(s2)

	var s3 *Student = &Student{"明明", 26}
	fmt.Println(*s3)

	var s4 *Student = &Student{
		Name: "娜娜",
		Age:  25,
	}
	fmt.Println(*s4)

}
