package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println(reType)
	fmt.Printf("reType类型:%v\n", reType)

	reValue := reflect.ValueOf(i)
	fmt.Println(reValue)
	fmt.Printf("reValue类型:%v\n", reValue)

	// num1 := 100
	num2 := 80 + reValue.Int()
	fmt.Println(num2)
	i2 := reValue.Interface()
	n := i2.(int)
	fmt.Println(n)
	n2 := n + 30
	fmt.Println(n2)
}
func main() {
	var num int = 100
	testReflect(num)
}
