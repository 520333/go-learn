package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("你好golang")
	}
	var temp int
	for {
		if temp >= 100 {
			break
		}
		temp++
		fmt.Println(temp)
	}

}
