package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	reValue := reflect.ValueOf(i)
	reValue.Elem().SetInt(20)

}
func main() {
	var num int = 100
	testReflect(&num)
	fmt.Println(num)
}
