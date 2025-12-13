package main

/*
import "fmt"

// 斐波那契数列
myfunc fibonacii(c chan int, quit chan int) {
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
myfunc main() {
	c := make(chan int)
	quit := make(chan int)
	go myfunc() {
		for i := 0; i < 50; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)

}
*/
