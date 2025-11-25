package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}
