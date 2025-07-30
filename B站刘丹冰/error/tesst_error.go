package main

import (
	"errors"
	"fmt"
)

func A() (int, error) { return 0, errors.New("this is a errors") }
func main() {
	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
