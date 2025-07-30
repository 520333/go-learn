package main

import (
	"fmt"
)

type Phone interface {
	Any
	Call(string)
}

type Camera interface {
	Take() string
}

type SmartPhone interface {
	Phone
	Camera
	Download(string) string
}

func ListSmartPhoneFunction(sp SmartPhone) {
	if v, ok := sp.(*Iphone); ok {
		v.Call("10086")
		fmt.Println("sp.Take()", v.Take())
		fmt.Println("sp.Download()", v.Download("www.baidu.com"))
	} else {
		fmt.Println("not Miphone PointerType")
	}
}

type MiPhone struct {
	Logo string
}

func (m *MiPhone) Call(phone string) {
	fmt.Println("Call to phone:", phone)
}

func (m *MiPhone) Take() string {
	return "logo.png"
}

func (m *MiPhone) Download(url string) string {
	return fmt.Sprintf("visit url:%s", url)
}

type Iphone struct {
	Logo string
}

func (m *Iphone) Call(phone string) {
	fmt.Println("Call to phone:", phone)
}

func (m *Iphone) Take() string {
	return "logo.png"
}

func (m *Iphone) Download(url string) string {
	return fmt.Sprintf("visit url:%s", url)
}

// 空接口
type Any interface{}

func GetAnyType(any interface{}) {
	switch t := any.(type) {
	case string:
		fmt.Println("any is string type")
	case int:
		fmt.Println("any is int type")
	case *MiPhone:
		fmt.Println("any is MiPhone type")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

}

type Queue []interface{}

func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func main() {
	// mp := new(MiPhone)
	// mp.Logo = "xiaomi"
	ip := new(Iphone)
	ip.Logo = "apple"
	ListSmartPhoneFunction(ip)

	var val Any
	val = 5
	fmt.Printf("val value:%v\n", val)
	GetAnyType(val)

	str := "ABCD"
	val = str
	fmt.Printf("val value:%v\n", val)
	GetAnyType(val)

	val = *ip
	fmt.Printf("val value:%v\n", val)
	GetAnyType(val)
	q := Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Push("abc")
	fmt.Println(q)

}
