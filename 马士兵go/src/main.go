package main

import "fmt"

func main() {
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'

	var s1 string = fmt.Sprintf("%d", n1)
	var s2 string = fmt.Sprintf("%f", n2)
	var s3 string = fmt.Sprintf("%t", n3)
	var s4 string = fmt.Sprintf("%c", n4)

	fmt.Printf("%T %v\n", s1, s1)
	fmt.Printf("%T %v\n", s2, s2)
	fmt.Printf("%T %q\n", s3, s3)
	fmt.Printf("%T %q\n", s4, s4)

}
