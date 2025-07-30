package main

import (
	"fmt"
	"reflect"
)

// 基本泛型
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// 类型约束
func sumNumbers[T int | float64](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func strs[T string](str T) T {
	fmt.Println(str)
	return str
}
func main() {
	PrintSlice([]int{1, 2, 3, 4, 5})
	PrintSlice([]float64{1.88, 3.14, 18.8})
	fmt.Println("==================")
	fmt.Println(sumNumbers([]float64{1.88, 3.14, 18.8}))
	fmt.Println(sumNumbers([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(reflect.TypeOf(sumNumbers([]int{1, 2})))
	strs("我是")
}
