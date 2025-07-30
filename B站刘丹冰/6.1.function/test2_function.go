package main

/*
import "fmt"

func add(a, b int) (sum int, err error) {
	sum = a + b
	return sum, err
}
func addany(item ...int) (sum int, err error) {
	for _, value := range item {
		sum += value
	}
	return sum, err
}
func addany2(a, b int) {
	fmt.Printf("sum is%d\r\n", a+b)
}

func calc(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("不是加法也不是减法")
		}
	}
}

func callBack(y int, f func(int, int)) {
	f(y, 2)
}


func main() {
	sum, err := add(1, 2)
	fmt.Println(sum, err)

	res, err := addany(1, 2, 3, 4)
	fmt.Println(res)
	// 函数赋值给变量
	funcVar := addany
	res1, _ := funcVar(5, 5, 5, 5)
	fmt.Println(res1)

	callBack(1, func(a, b int) {
		fmt.Printf("total is:%d\r\n", a+b)
	})


}
*/
