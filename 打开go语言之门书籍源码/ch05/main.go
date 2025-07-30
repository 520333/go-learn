package main

import (
	"errors"
	"fmt"
)

// 传入2个参数,返回一个整数
func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}
func sum2(params ...int) int {
	var sum int
	for _, i := range params {
		sum += i
	}
	return sum
}

// 匿名函数
var sum1 = func(a, b int) int {
	return a + b
}

// 闭包函数
func autoIncrement() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 递归
func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		fmt.Printf("当前数值:%d 当前计算结果:%d\r\n", n, result)
		return
	} else {
		return 1
	}
}

type Age uint

func (age Age) string() {
	fmt.Printf("this age is %d\r\n", age)
}
func (age *Age) modify() {
	*age = Age(30)
}
func main() {
	result, err := sum(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(sum2(4, 2, 5, 10))
	fmt.Println(sum1(1, 2))
	fmt.Println("=============")
	next := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(next())
	}
	fmt.Println("=============")
	for i := 0; i < 5; i++ {
		fmt.Println(next())
	}

	age := Age(20)
	age.string()
	(&age).modify()
	age.string()

	i := 5
	res := factorial(5)
	fmt.Println(i, res)
	age1 := Age(40)
	sm := Age.string
	sm(age1)
}
