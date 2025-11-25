package main

import "fmt"

func main() {
	var score int = 91
	if score >= 90 {
		fmt.Println(score, "A")
	} else if score <= 90 && score >= 80 {
		fmt.Println(score, "B")
	} else if score <= 80 && score >= 70 {
		fmt.Println(score, "C")
	} else if score <= 70 && score >= 60 {
		fmt.Println(score, "D")
	} else if score < 60 {
		fmt.Println(score, "E")
	}
}
