package main

import "fmt"

func main() {
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21
	var sum, avg int
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg = sum / len(scores)
	fmt.Println("总分：", sum)
	fmt.Println("平均分：", avg)
	fmt.Println(len(scores), cap(scores))
}
