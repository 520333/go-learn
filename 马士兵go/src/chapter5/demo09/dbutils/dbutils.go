package dbutils

import "fmt"

func GetConnection() {
	fmt.Println("执行dbutils包下的getConnection函数")
}

func ExeChange(num1, num2 *int) {
	var temp = *num1
	*num1 = *num2
	*num2 = temp
}

func init() {
	fmt.Println("init...")
}
