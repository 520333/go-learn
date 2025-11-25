package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 2)
	fmt.Println(s1)

	var n2 float32 = 4.29
	var s2 string = strconv.FormatFloat(float64(n2), 'f', 9, 64)
	fmt.Println(s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3=%q \n", s3)
}
