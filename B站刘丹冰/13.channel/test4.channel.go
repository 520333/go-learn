package main

/*
import "fmt"

// 斐波那契数列
func fibonacii(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)

}
*/
