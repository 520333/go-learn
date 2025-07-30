package main

/*
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", u)
}

func DoFileAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputType = ", inputType.Name())
	fmt.Println("inputValue = ", inputValue)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("方法: %s 类型: %v\n", method.Name, method.Type)
	}
}
func main() {
	user := User{1, "dawn", 20}
	user.Call()
	DoFileAndMethod(user)
}
*/
