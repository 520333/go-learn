package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) test1() {
	(*p).Name = "蟹老板"
	fmt.Println((*p).Name)
	fmt.Printf("p的地址:%p \n", &p)
}

func (p *Person) test2() {
	p.Name = "章鱼哥"
	fmt.Println(p.Name)
	fmt.Printf("p的地址:%p \n", &p)
}

func main() {
	var p Person
	p.Name = "派大星"
	fmt.Printf("p的地址:%p \n", &p)
	(&p).test1()
	fmt.Println(p.Name)

	p.test2()
	fmt.Println(p.Name)

}
