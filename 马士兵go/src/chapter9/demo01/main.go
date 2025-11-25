package main

import "fmt"

func main() {
	var a map[int]string //方式1
	a = make(map[int]string, 10)

	a[20095452] = "张三"
	a[20095387] = "李四"
	a[20097291] = "王五"
	fmt.Println(a)
	for i := 0; i < len(a); i++ {

		fmt.Printf("map[%v]= \t\n", i)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Printf("%p\n", &a)

	b := make(map[int]string, 10) //方式2
	b[200912] = "张三"
	b[200913] = "王五"
	fmt.Println(b)

	c := map[int]string{
		20095452: "张三",
		20095429: "王五",
	}
	fmt.Println(c)
}
