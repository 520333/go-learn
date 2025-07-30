package main

import "fmt"

type Person struct {
	name string
}

// 接收者
func (p *Person) Sayhello() {

}

// 通过指针交换2个值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func changeName(p *Person) {
	p.name = "阿宝"
}
func main() {
	p := Person{
		name: "dawn",
	}
	var pi *Person = &p
	// changeName(&p)
	// fmt.Println(p.name)
	fmt.Printf("%p", pi)

	// 指针定义
	var po *Person
	po = &p
	fmt.Println(po)
	// var a int = 10
	// b := &a
	//var b *int

	p2 := &Person{} // 指针第一种初始化方式
	fmt.Println(p2.name)
	var emptyPerson Person

	p3 := &emptyPerson // 指针第二种初始化方式
	fmt.Println(p3.name)

	var p4 *Person = new(Person) // 指针第三种初始化方式 推荐方式
	fmt.Println(p4.name)
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

}
