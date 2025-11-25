package main

import "fmt"

type Integer int

func (i Integer) print() {
	i = 30
	fmt.Println("i =", i)
}
func (i *Integer) change() {
	*i = 40
	fmt.Println("i =", *i)
}
func main() {
	var i Integer = 20
	i.print()
	i.change()
	fmt.Println(i)

}
