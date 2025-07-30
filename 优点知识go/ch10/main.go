package main

import (
	"fmt"
	"math"
	"os"
)

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("not support operate: %s", op)
	}

}

func swap(a, b int) (int, int) {
	return b, a
}

func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}
func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type greeting func(name string) string

func say(g greeting, word string) {
	fmt.Println(g(word))
}
func english(name string) string {
	return "hello," + name
}
func french(name string) string {
	return "Bonjour," + name
}
func china(name string) string {
	return "你好，" + name
}

func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
func main() {
	fmt.Println(operate(10, 5, "x"))
	fmt.Println(swap(10, 5))

	file, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println("打开文件错误检查是否文件存在")
	} else {
		fmt.Println(file)
	}

	fmt.Println(compute(powInt, 10, 5))

	say(french, "world")

	fmt.Println(sum(1, 2))
}
