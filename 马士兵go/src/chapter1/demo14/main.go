package main

import "fmt"

func main() {
	var age int = 18
	fmt.Println(&age)
	age2 := &age

	fmt.Println(*age2)

	var prt *int = &age
	fmt.Println(&prt)

}
