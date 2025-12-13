package main

import (
	"fmt"
	"time"
)

func getData(i int, c chan int) {
	for n := range c {
		// n := <-c
		// if n, ok := <-c; ok {
		// 	fmt.Printf("get data %d from channel %d\n", n, i)
		// } else {
		// 	break
		// }
		fmt.Printf("get data %d from channel %d\n", n, i)

	}
}
func createChannel(i int) chan<- int {
	c := make(chan int)
	go getData(i, c)
	return c
}
func chanelDemo() {
	// c := make(chan int)
	// go myfunc() {
	// 	for {
	// 		n := <-c
	// 		fmt.Println(n)
	// 	}
	// }()
	// c <- 10
	// c <- 20
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createChannel(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i + 100
	}
	for i := 0; i < 10; i++ {
		close(channels[i])
	}

	// c := createChannel()
	// go getData(c)
	// c <- 10
	// c <- 20
	time.Sleep(time.Millisecond)
}
func main() {
	chanelDemo()

}
