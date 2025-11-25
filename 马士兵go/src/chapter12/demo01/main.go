package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello golang " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	go test()
	for i := 1; i <= 5; i++ {
		fmt.Println("main golang " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
	}
}
