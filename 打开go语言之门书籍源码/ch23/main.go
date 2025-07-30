package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	int | float32 | float64 | int32
}

func min[V cmp.Ordered](a, b V) V {
	if a < b {
		return a
	} else {
		return b
	}
}

func map1(s []int, f func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type MySlice []int

func (s MySlice) map1(f func(int) int) MySlice {
	result := make(MySlice, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type GSlice[T any] []T

func (s GSlice[T]) map1(f func(T) T) GSlice[T] {
	result := make(GSlice[T], len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}
func (s GSlice[T]) reduce(f func(previousValue T, currentValue T) T) T {
	var result T
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

func (s GSlice[T]) filter(f func(T) bool) GSlice[T] {
	result := GSlice[T]{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	fmt.Println(min(12, 2))
	fmt.Println(min[float32](3.14, 1))
	fmt.Println(min[int32](3, 1))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(map1(s, func(v int) int {
		return v * 2
	}))

	s1 := MySlice{1, 2, 3, 4, 5}
	fmt.Println(s1.map1(func(v int) int {
		return v * 2
	}))
	gs1 := GSlice[int]{1, 3, 5, 7, 9}
	fmt.Println(gs1.map1(func(v int) int {
		return v * 2
	}))
	gss1 := GSlice[string]{"1", "3", "5", "7", "9"}
	fmt.Println(gss1.map1(func(v string) string {
		return v + "阿宝"
	}))
	sum := gs1.map1(func(v int) int {
		return v * 2
	}).reduce(func(previousValue, currentValue int) int {
		return previousValue + currentValue
	})
	fmt.Println(sum)
	fmt.Println(gs1.filter(func(v int) bool {
		return v > 3
	}))

	sum1 := gs1.filter(func(v int) bool {
		return v > 3
	}).map1(func(v int) int {
		return v * 2
	}).reduce(func(preV int, curV int) int {
		return preV + curV
	})
	fmt.Println("filter map reduce >> sum:", sum1)
}
