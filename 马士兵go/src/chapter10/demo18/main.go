package main

import (
	"fmt"
)

type SayHello interface {
	SayHello()
}

type Chinese struct {
	Name string
}

func (c Chinese) SayHello() {
	fmt.Println("你好!!!")
}
func (c Chinese) niuYangGe() {
	fmt.Println("东北文化-扭秧歌!!!")
}

type American struct {
	Name string
}

func (a American) SayHello() {
	fmt.Println("Hello!!!")
}
func (a American) disco() {
	fmt.Println("野狼disco")
}

func greet(s SayHello) {
	s.SayHello()
	//类型断言
	// if ch, flag := s.(Chinese); flag {
	// 	ch.niuYangGe()
	// } else {
	// 	fmt.Println("美国人不会扭秧歌!!!")
	// }

	// switch s.(type) {
	// case Chinese:
	// 	ch := s.(Chinese)
	// 	ch.niuYangGe()
	// case American:
	// 	am := s.(American)
	// 	am.disco()
	// }
	switch v := s.(type) {
	case Chinese:
		v.niuYangGe()
	case American:
		v.disco()
	}
	fmt.Println("打招呼")
}

type intger int

func (i intger) SayHello() {
	fmt.Println("say hi +", i)
}
func main() {
	c := Chinese{}
	greet(c)
	a := American{}
	greet(a)

}
