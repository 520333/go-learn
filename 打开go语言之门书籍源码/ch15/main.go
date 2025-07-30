package main

import "fmt"

func main() {
	var s string
	s = "张三"
	fmt.Println(s)
	fmt.Printf("%p\n", &s)

	var sp *string
	sp = new(string)
	*sp = "海绵宝宝"
	fmt.Println(*sp)

	pp := NewPerson("海绵宝宝", 20)
	fmt.Println("name为", pp.name, "age为：", pp.age)

	m := map[string]int{"张三": 18}
	fmt.Println(m)
}
func NewPerson(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p
}

type person struct {
	name string
	age  int
}
