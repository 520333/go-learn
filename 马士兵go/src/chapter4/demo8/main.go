package main

import "fmt"

func main() {
	var str string = "hello golang你好"
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c\n", str[i])
	//}
	for i, v := range str {
		fmt.Printf("%d %c\n", i, v)
	}
}
