package main

import "fmt"

func tryPanic() {

	fmt.Println("first line in tryPanic function")
	// panic("call a panic")
	a := 0
	b := 100
	r := b / a
	fmt.Println(r)
	fmt.Println("second line in tryPanic function")
}
func tryPanic2() {

	fmt.Println("first line in tryPanic function")
	// panic("call a panic")
	a := 0
	b := 100
	r := b / a
	fmt.Println(r)
	fmt.Println("second line in tryPanic function")
}
func protect(g ...func()) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("catch panic in recover function: %s\n", e)
		}
	}()
	for _, f := range g {
		f()
	}
}

type panicFunc func()

func main() {
	fmt.Println("start call tryPanic MAIN")
	// tryPanic()
	fmt.Println("end call tryPanic MAIN")
	var a = new(panicFunc)
	protect(tryPanic2, tryPanic, *a)
}
