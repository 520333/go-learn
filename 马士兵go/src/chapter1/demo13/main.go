package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b=%v \n", b)

	var s2 string = "19"
	var num int64
	num, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num=%v %T \n", num, num)

	var s3 string = "3.14"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("f1 %v \n", f1)

	var s4 string = "golang"
	var b1 bool
	b1, _ = strconv.ParseBool(s4)
	fmt.Printf("b1 %v \n", b1)
}
