package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("hello golang" + strconv.Itoa(n))
		}(i)
	}
	wg.Wait()
}
