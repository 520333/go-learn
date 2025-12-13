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

myfunc (u User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", u)
}

myfunc main() {
	user := User{1, "DAWN", 20}
	DoFileAndMethod(user)
}
myfunc DoFileAndMethod(input interface{}) {
	// 获取input type
	inputType := reflect.TypeOf(input)
	// 获取input value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputType = ", inputType.Name())
	fmt.Println("inputValue = ", inputValue)

	// 获取里面的字段在通过type获取里面的方法
	//1.获取interface的reflect.Type。通过Type得到NumField
	//2.得到每个Field 数据类型
	//3.通过field有一个interface()方法得到对应的value
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
*/
