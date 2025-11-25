package main

import "fmt"

type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
	a int
}

type C struct {
	A
	B
	int
}
type C1 struct {
	*A
	*B
	int
}
type D struct {
	a int
	b string
	c B
}

func main() {
	c := C{A{10, "aaa"}, B{20, "bbb", 50}, 888}
	fmt.Println(c.B.a)
	fmt.Println(c.int)

	c1 := C1{&A{10, "c1"}, &B{20, "B1", 50}, 888}
	fmt.Println(*c1.A)
	fmt.Println(*c1.B)

	var d D = D{10, "ooo", B{10, "b", 10}}
	fmt.Println(d.c.d)
}
