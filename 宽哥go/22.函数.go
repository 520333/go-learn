package main

import (
	"fmt"
	"strings"
)

// func 函数名(参数1,参数2,类型)(返回值1,返回值2 类型){代码块/函数体}
func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func qiuhe(i1, i2 int) (sum int) {
	sum = i1 + i2
	return sum
}
func paixu(i1, i2 int) (min, max int) {
	if i1 > i2 {
		max = i1
		min = i2
	} else {
		max = i2
		min = i1
	}
	return
}

func hasName(s ...string) string {

	fmt.Println("接收不定长度的函数:", s)
	m := strings.Join(s, "-")
	fmt.Println("拼接后的字符串:", m)
	return m
}

// func main() {
// 	res := max(31, 2)
// 	fmt.Println("最大值:", res) //stdout 最大值:31
// 	res = qiuhe(5, 11)
// 	fmt.Println("求和:", res) //stdout 求和: 16
// 	min1, max1 := paixu(52, 1121)
// 	fmt.Println("从小到大排序:", min1, max1) //stdout 从小到大排序: 52 1121

// 	hasName("chuang", "dawn")            //stdout: 接收不定长度的函数: [chuang dawn]
// 	hasName("kubernetes", "statefulset") //stdout: 拼接后的字符串: kubernetes-statefulset

// }
