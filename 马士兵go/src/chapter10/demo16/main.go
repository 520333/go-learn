package main

import "fmt"

type AInterFace interface {
	a()
}
type BInterFace interface {
	b()
}
type Stu struct {
}

func (s Stu) a() {
	fmt.Println("aaaa")
}
func (s Stu) b() {
	fmt.Println("bbbb")
}
func main() {
	var s Stu
	var a AInterFace = s
	var b BInterFace = s
	a.a()
	b.b()
}
