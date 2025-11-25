package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan值:%v\n", intChan)
	intChan <- 10
	num := 20
	intChan <- num
	close(intChan)

	fmt.Println(len(intChan), cap(intChan))
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

}
