package main

import "fmt"

func test(num *int) {
	*num = 30
	fmt.Println("test----", *num)
	fmt.Printf("test----%p\n", num)

}
func main() {
	var num = 10
	test(&num)
	fmt.Println("main----", num)
	fmt.Printf("main----%p\n", &num)
}
