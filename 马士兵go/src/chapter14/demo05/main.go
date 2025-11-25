package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) CPrint() {
	fmt.Println("调用了print()方法")
	fmt.Printf("学生的名字:%v\n", s.Name)
}
func (s Student) AGetSum(n1, n2 int) int {
	fmt.Println("调用了AGetSum()方法")
	return n1 + n2
}
func (s *Student) BSet(name string, age int) {
	s.Name = name
	s.Age = age
	fmt.Println("调用了BSet()方法")

}
func TestStudent(i interface{}) {
	val := reflect.ValueOf(i)
	fmt.Println(val)
	n1 := val.Elem().NumField()
	fmt.Printf("结构体字段数:%v\n", n1)
	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个字段的值是:%v\n", i+1, val.Elem().Field(i))
	}
	n2 := val.NumMethod()
	fmt.Println(n2)

	val.Method(2).Call(nil)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	result := val.Method(0).Call(params)
	fmt.Println(result[0].Int())

	params1 := []reflect.Value{reflect.ValueOf("海绵宝宝"), reflect.ValueOf(25)}
	val.MethodByName("BSet").Call(params1)

}
func main() {
	var s Student = Student{"派大星", 18}
	TestStudent(&s)
	fmt.Println("最终 s =", s)

}
