package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var i int = 10
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	var (
		j int = 0
		k int = 1
	)
	var (
		p = 0
		q = 1
	)
	fmt.Println(j, k, p, q)
	var fl32 float32 = 2.11
	var fl64 float64 = 10.224
	fmt.Println("fl32 is", fl32, "fl64 is ", fl64)
	pi := &i
	fmt.Println(pi)
	i = 20
	fmt.Println(*pi)

	const name = "dawn"
	fmt.Println(name)
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
	// i2s := strconv.Itoa(i)
	// s2i, err := strconv.Atoi(i2s)
	// fmt.Printf("i2s %s %T\r\n", i2s, i2s)
	// fmt.Printf("s2i %d %T  %v\r\n", s2i, s2i, err)
	i2f := float64(i)
	f2i := int(fl32)
	fmt.Println(i2f, reflect.TypeOf(i2f), f2i, reflect.TypeOf(f2i))
	s1 := "Hello,world!"
	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.ToUpper(s1))

}
