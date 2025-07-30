package main

import (
	"fmt"
	"sync"
)

type MyInt int
type Number interface {
	~int | ~int32 | float32 | float64 | string
}

func sum[T Number](x, y T) T {
	return x + y
}

func MapValues[T Number](values []T, myFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := myFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

type CustomData interface {
	[]byte | []rune
}
type User[T comparable] struct {
	Id   int
	Name string
	Data T
}

type CustomMap[T comparable, V int | string] map[T]V

func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel Closed")
			return
		}
		fmt.Printf("Reader %d Received %d\n", id, val)
	}

}
func main() {
	// fmt.Println("使用泛型:", sum("3123", "23.123"))
	// fmt.Println("使用泛型:", sum(MyInt(10), MyInt(20)))

	// result := MapValues([]float64{1.1, 2.2, 3.3}, func(i float64) float64 {
	// 	return i * 2
	// })
	// fmt.Println(result)
	// u := User[string]{
	// 	Id:   0,
	// 	Name: "阿宝",
	// 	Data: "aa",
	// }
	// fmt.Println(u)

	// m := make(CustomMap[int, string])
	// m[3] = "333"
	// fmt.Println(m)

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(4)
	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
