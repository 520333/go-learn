package main

import "fmt"

//interface是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...", arg)
	fmt.Printf("myFunc data: %v  myFunc type is %T\n", arg, arg)
	fmt.Println("==============================")
	//interface会启动类型断言机制 判断是什么类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Printf("arg type is %T, value = %v\n", value, value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	myFunc(book)
	myFunc(100)

	aa := map[int]string{
		100: "java",
		200: "java",
		300: "java",
	}
	myFunc(aa)
	myFunc(3.14)
	myFunc("我是")

}
