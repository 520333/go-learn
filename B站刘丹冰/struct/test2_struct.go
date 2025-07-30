package main

import "fmt"

// 结构体嵌套
type Person struct {
	name string
	age  int
}
type Student1 struct {
	p     Person // 具名嵌套
	score float32
}
type Student2 struct {
	Person // 匿名嵌套
	score  float32
	name   string
}
type Student3 struct {
	p       Person // 具名嵌套
	score   float32
	address struct {
		city string
	}
}

func (p Person) print() { //值传递
	fmt.Printf("name:%s, age:%d\r\n", p.name, p.age)
}
func (p *Person) printp() { //引用传递
	p.age = 2000
	fmt.Printf("name:%s, age:%d\r\n", p.name, p.age)
}
func main() {
	s := Student1{
		Person{
			"dawn",
			18,
		},
		95.6,
	}
	fmt.Println(s.p.age)
	s2 := Student2{
		Person{
			"chuang",
			20,
		},
		100,
		"aaaa", // 外部优先级高
	}
	s2.name = "bbbb"
	fmt.Println(s2)
	fmt.Println(s2.age, s2.name)
	s2.print()
	p2 := Person{
		"yoda", 1000,
	}
	p2.print()
	p2.print()
}
