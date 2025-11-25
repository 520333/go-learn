package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 1; i < 100; i++ {
		intChan <- i
	}
	close(intChan)
	// fmt.Println(len(intChan))
	for v := range intChan {
		fmt.Println(v)
	}
}
