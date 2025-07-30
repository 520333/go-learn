package main

import (
	"fmt"
	"math"
)

func constDemo() {
	const s string = "Hello"
	const a, b = 3, 4
	fmt.Println(s, a, b)
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}
func enumDemo() {
	const (
		Sunday = iota + 1
		Monday = iota * 2
		Tuesday
		Wednesday
		Thursday
		Friday
		Staurday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday)
}
func main() {
	constDemo()
	enumDemo()

}
