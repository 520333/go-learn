package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	close(intChan)
	// intChan <- 30
	num := <-intChan
	fmt.Println(num)

}
