package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)

	reValue := reflect.ValueOf(i)
	k1 := reType.Kind()
	fmt.Println(k1)
	k2 := reValue.Kind()
	fmt.Println(k2)

	i2 := reValue.Interface()
	n, flag := i2.(Student)
	if flag {
		fmt.Println(n.Name, n.Age)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {
	var s Student = Student{Name: "派大星", Age: 18}
	testReflect(s)
	var num = 100
	testReflect(num)
}
