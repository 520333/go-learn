package main

import "fmt"

type SayHello interface {
	SayHello()
}

type Chinese struct {
	Name string
}

func (c *Chinese) SayHello() {
	fmt.Println("你好!!!")
}

type American struct {
	Name string
}

func (a American) SayHello() {
	fmt.Println("Hello!!!")
}
func greet(s SayHello) {
	s.SayHello()
}

type intger int

func (i intger) SayHello() {
	fmt.Println("say hi +", i)
}
func main() {
	var arr [3]SayHello
	arr[0] = American{"rose"}
	arr[1] = &Chinese{"丽丽"}
	arr[2] = &Chinese{"菲菲"}
	fmt.Println(arr)
	var c Chinese = Chinese{}
	var a American = American{}
	greet(&c)
	greet(a)

	var s SayHello
	s = &Chinese{}
	s.SayHello()
	s = American{}
	s.SayHello()

	var i intger = 10
	var s1 SayHello = i
	s1.SayHello()

}
