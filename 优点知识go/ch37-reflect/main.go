package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" ok:"a"`
	Age  int    `json:"age" ok:"b"`
}

func main() {
	// reflectBasic()
	// reflectLoopStruct()
	// reflectChangeFiledVale()
	reflectDynamicCallMethod()
	reflectStructTag()
}
func (u User) String(prefix string) {
	fmt.Printf("Prefix=%s Name=%s Age=%d\n", prefix, u.Name, u.Age)
}
func (u User) Print() {
	fmt.Println("hello reflect!!!")
}
func reflectBasic() {
	u := User{"dawn", 30}
	t := reflect.TypeOf(u)
	fmt.Printf("TypeOf(u)=%v\n", t)
	v := reflect.ValueOf(u)
	t0 := v.Type()
	fmt.Printf("ValueOf(u)=%v reflect.Type(u)=%v\n", v, t0)

	fmt.Printf("%T %v \n", u, u)
	fmt.Println("=================")

	u1 := v.Interface().(User)
	fmt.Println(u1, reflect.TypeOf(u1))

	fmt.Println("=================")
	fmt.Println(t0.Kind())
}

func reflectLoopStruct() {
	u := User{"张三", 19}
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("fileIndex:%d fileName:%s\n", f.Index, f.Name)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("MethodIndex:%d MethodName:%s\n", m.Index, m.Name)
	}
}

func reflectChangeFiledVale() {
	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println(x)
}

func reflectDynamicCallMethod() {
	u := User{"宝哥云", 20}
	v := reflect.ValueOf(u)

	printM := v.MethodByName("String1")
	if printM.IsValid() {
		args := []reflect.Value{reflect.ValueOf("OK")}
		fmt.Println(printM.Call(args))
	}

}
func reflectStructTag() {
	var u User
	h := `{"name": "dawn","age": 20}`
	if err := json.Unmarshal([]byte(h), &u); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(u, u.Name, u.Age)
	}
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag, f.Tag.Get("ok"))
	}

}
