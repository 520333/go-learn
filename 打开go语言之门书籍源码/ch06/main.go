package main

import "fmt"

type person struct {
	name string
	age  uint
	address
}
type address struct {
	province string
	city     string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// 工厂函数
func NewPerson(name string) *person {
	return &person{name: name}
}

type errorString struct {
	s string
}

func New(text string) error {
	return &errorString{text}
}
func (e *errorString) Error() string {
	return e.s
}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

func (p *person) Walk() { fmt.Printf("%s能走\n", p.name) }
func (p *person) Run()  { fmt.Printf("%s能跑\n", p.name) }

type WalkRun interface {
	Walk()
	Run()
}

func main() {
	p := person{name: "dawn", age: 20, address: address{province: "福建省", city: "泉州市"}}
	fmt.Println(p.name, p.age)
	printString(&p)
	printString(p.address)
	fmt.Println(NewPerson("宝哥"))
	fmt.Println(New("string"))

	p1 := NewPerson("张三")
	fmt.Println(p1)
	var s fmt.Stringer
	s = p1

	p2, err := s.(*person)
	if err {
		fmt.Println(*p2)
	} else {
		fmt.Println("s不是一个person")
	}
	fmt.Println(p2)
	WalkRun.Run(&p)
	WalkRun.Walk(&p)
}
