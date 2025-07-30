package main

import "fmt"

func deferReturn() (ret int) {
	defer func() {
		ret++
		fmt.Println("return之后打印")
	}()
	fmt.Println("return之前打印")
	return 10

}
func main() {
	// 连接数据库 最后无论如何都要关闭数据库
	ret := deferReturn()
	fmt.Printf("ret=%d\r\n", ret)
}
