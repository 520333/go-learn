package main

import "fmt"

type CInterFace interface {
	c()
}
type BInterFace interface {
	b()
}
type AInterFace interface {
	BInterFace
	CInterFace
	a()
}
type Stu struct {
}

func (s Stu) a() {
	fmt.Println("a function.")
}
func (s Stu) b() {
	fmt.Println("b function.")
}
func (s Stu) c() {
	fmt.Println("c function.")
}

type E interface{}

func main() {
	var s Stu
	var a AInterFace = Stu{}
	a.a()
	a.b()
	a.c()
	var e E = s
	fmt.Println(e)
	var e2 interface{} = e
	fmt.Println(e2)
	var num float64 = 9.3
	var e3 interface{} = num
	fmt.Println(e3)
}
