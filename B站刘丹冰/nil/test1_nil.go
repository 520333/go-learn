package main

import "fmt"

// 不用类型的零值不一样
/*bool false
numbers 0
string ""
pointer nil
slice nil
map nil
channel、interface、function nil
struct默认不是nil 零值取决于里面的字段类型
*/
type Person struct {
	name string
	age  int
	f    *int
}

func main() {
	p1 := Person{
		name: "dawn",
		age:  19,
	}
	p2 := Person{
		name: "abao",
		age:  18,
	}
	if p1 == p2 {
		fmt.Println("yes")
	}
	fmt.Println()

	var ps []Person
	var ps2 = make([]Person, 0)
	if ps == nil {
		fmt.Println("nil slice")
	}
	if ps2 == nil {
		fmt.Println("ps2 is nil")
	}
	var m map[string]string             //nil Map
	var m2 = make(map[string]string, 0) // emptyMap
	if m == nil {
		fmt.Println("nil map")
	}
	if m2 == nil {
		fmt.Println("empty map")
	}
}
