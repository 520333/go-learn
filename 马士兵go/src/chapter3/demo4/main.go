package main

import "fmt"

func main() {
	//与逻辑 &&
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//或逻辑 ||
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//非逻辑 !
	fmt.Println(!true)
	fmt.Println(!false)

}
