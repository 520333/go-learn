package main

import "fmt"

// 实现数字阶乘: 5*4*3*2*1
func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		fmt.Println("当前数值:", n, "当前计算结果:", result)
		return
	} else {
		return 1
	}

}

// func main() {
// 	//递归函数:函数自己调用自己，直到达到条件后才能结束返回结果
// 	//n!=n*(n-1)*(n-2)***1
// 	i := 5
// 	res := factorial(5)
// 	fmt.Printf("%d的阶乘结果:%d", i, res) //stdout: 5的阶乘结果:120
// }
