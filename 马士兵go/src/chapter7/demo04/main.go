package main

import "fmt"

func main() {
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学员成绩:", i+1)
		fmt.Scanln(&scores[i])
	}
	var sum, avg int
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg = sum / len(scores)
	fmt.Println("总分：", sum)
	fmt.Println("平均分：", avg)

	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为：%d \n", i+1, scores[i])
	}
	fmt.Println("-----------")
	for i, v := range scores {
		fmt.Printf("第%d个学生的成绩为：%d \n", i+1, v)
	}
}
