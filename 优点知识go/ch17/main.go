package main

import (
	"fmt"
	"learn/ch17/book"
	"strconv"
)

type Integer int

func (it Integer) String() string {
	return strconv.Itoa(int(it))
}

func main() {
	bk := book.NewBook(1, "java", "dawn", "数据类型")
	fmt.Println(bk.String())
	book.RefTag(*bk, 0)
	book.InitTechBook()

	it := Integer(100)
	fmt.Printf("it=%s type=%T", it, it)
}
