package main

import "fmt"

func main() {
	b := make(map[int]string)
	b[20095452] = "张三"
	b[20095387] = "李四"

	b[20095452] = "王五"
	delete(b, 20095452)
	fmt.Println(b)

	v, f := b[20095387]
	fmt.Println(v, f)

	b = map[int]string{}
	fmt.Println(b)
}
