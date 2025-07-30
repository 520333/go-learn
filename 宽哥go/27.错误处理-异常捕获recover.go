package main

// func printSliceData(s []string) {
// 	defer func() { //匿名函数
// 		fmt.Println("程序执行失败，捕获异常")
// 		if err := recover(); err != nil {
// 			//recover用来捕获panic的报错 尝试恢复，防止程序异常退出
// 			fmt.Println("捕获到一个错误:", err)
// 		}
// 	}()
// 	fmt.Println("切片的内容:", s)
// 	fmt.Println("切片的第三个值是:", s[2])
// }
// func main() {
// 	s := []string{"a", "b"}
// 	printSliceData(s)
// }
