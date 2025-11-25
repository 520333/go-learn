package main

import "fmt"

func main() {
	fmt.Println("hello golang1")
	fmt.Println("hello golang2")
	fmt.Println("hello golang3")
	goto label1
	fmt.Println("hello golang4")
	fmt.Println("hello golang5")
	fmt.Println("hello golang6")
label1:
	fmt.Println("hello golang7")
}
