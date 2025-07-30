package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp = main::a
	*pa = *pb  // main:a = main:b
	*pb = temp // main:b = main:a

}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("a 的内存地址:", &a)
	fmt.Println("b 的内存地址:", &b)

	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)
	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)

}
