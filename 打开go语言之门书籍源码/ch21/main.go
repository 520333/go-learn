package main

import (
	"fmt"
	"os"
)

func main() {
	mk := os.Mkdir("tmp", 0666)
	if mk != nil {
		fmt.Println(mk.Error())
	}
}
func newString() string {
	s := new(string)
	*s = "宝哥无情"
	return *s
}
