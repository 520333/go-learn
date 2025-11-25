package main

import "fmt"

func main() {
	var score int = 71
	switch score / 10 {
	case 10:
		fmt.Println(score, "A")
	case 9:
		fmt.Println(score, "B")
	case 8:
		fmt.Println(score, "C")
	case 7:
		fmt.Println(score, "D")
	case 1, 2, 3, 4, 5, 6:
		fmt.Println(score, "E")
	}
	var a int = 2
	switch {
	case a == 1:
		fmt.Println("aaa")
	case a == 2:
		fmt.Println("bbb")
	}
	switch b := 20; {
	case b > 10:
		fmt.Println("大于10")
	}

	var sum int
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
