package main

import "fmt"

func main() {
	var intChan chan<- int //写channel
	intChan = make(chan<- int, 3)
	for i := 0; i < 3; i++ {
		intChan <- i
	}
	fmt.Println(intChan)

	var intChan3 <-chan int
	if intChan3 != nil {
		num1 := <-intChan3
		fmt.Println(num1)
	}
}
