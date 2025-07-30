package main

//defer会在return之后执行
/*import "fmt"

func a() {
	fmt.Println("A")
}
func b() {
	fmt.Println("B")
}
func c() {
	fmt.Println("C")
}
func main() {
	defer fmt.Println("main end1") // 顺序：先进栈后出栈
	defer fmt.Println("main end2") // 顺序：后进栈先出栈

	fmt.Println("main:hello go1")
	fmt.Println("main:hello go2")

	defer a()
	defer b()
	defer c()
}
*/
