package main

import (
	"fmt"
	"reflect"
)

func numOperations(a, b int) {
	fmt.Printf("%d + %d= %d \n", a, b, a+b)
	fmt.Printf("%d - %d= %d \n", a, b, a-b)
	fmt.Printf("%d * %d= %d \n", a, b, a*b)
	fmt.Printf("%d / %d= %f \n", a, b, float64(a)/float64(b))
	fmt.Printf("%d 取余 %d= %d \n", a, b, a%b)
}

func stringOperations(a, b string) {
	fmt.Printf("a和b拼接后：%s \n", a+b)
	ab := a + b
	fmt.Printf("ab: %s,类型是：%s", ab, reflect.TypeOf(ab))
}
func stringSprintf(firstName, secondName string) {
	// fullName := secondName + firstName
	fullName := fmt.Sprintf("%s%s", secondName, firstName)
	fmt.Println("全名：", fullName)
}

// func main() {
// 	// 数值运算
// 	// num1 := 1
// 	// num2 := 2
// 	// fmt.Println(num1 + num2)
// 	numOperations(1, 5)

// 	name1 := "chuang"
// 	name2 := "宝哥"
// 	stringOperations(name1, name2)
// 	stringOperations("1", "2")
// 	stringSprintf("四", "李")
// 	p1 := 9
// 	p1++
// 	fmt.Println("p1自增后的值：", p1)
// 	p1--
// 	fmt.Println("p1自减后的值：", p1)

// }
