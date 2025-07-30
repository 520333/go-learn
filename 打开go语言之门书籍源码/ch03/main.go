package main

import "fmt"

func main() {
	i := 6
	if i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("i>5 && i<10")
	} else {
		fmt.Println("i<=10")
	}
	if i := 10; i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("i>5 && i<10")
	} else {
		fmt.Println("i<=10")
	}

	switch i = 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("i>5 && i<=10")
	default:
		fmt.Println("i<=5")
	}

	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}
	switch 2 > 1 {
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")
	}

	sum := 0
	i = 1
	// for i = 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// for i <= 100 {
	// 	sum += i
	// 	i++
	// 	if i > 100 {
	// 	}
	// }
	// fmt.Println(sum)
	// for {
	// 	sum += i
	// 	i++
	// 	if i > 100 {
	// 		break
	// 	}
	// }
	fmt.Println(sum)

	for j := 1; j < 100; j++ {
		if j%2 != 0 {
			continue
		}
		sum += j
	}
	fmt.Println(sum)
}
