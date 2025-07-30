package main

import "fmt"

func main() {
	const PI float32 = 3.14 // 常量定义了可以不使用 常量名尽量使用大写
	const (
		ERR1, ERR2 = 1 + iota, 2 + iota //iota=0
		ERR3, ERR4                      //iota=1
		ERR5, ERR6 = iota * 2, iota * 3 //iota=2 ERR5=4 ERR6=6
		ERR7, ERR8                      //iota=3 ERR7=6 ERR8=9
		ERR9       = "ok"               //iota=4
		ERR10      = 100                //iota=5
		ERR11      = iota               //iota=6
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR4)
	fmt.Println(ERR5, ERR6, ERR7, ERR8)
	fmt.Println(ERR9, ERR10, ERR11)
}
