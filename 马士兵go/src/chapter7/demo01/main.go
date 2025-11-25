package main

import "fmt"

func main() {
	score1 := 95
	score2 := 91
	score3 := 39
	score4 := 60
	score5 := 21
	sum := score1 + score2 + score3 + score4 + score5
	avg := (score1 + score2 + score3 + score4 + score5) / 5
	fmt.Println("总分：", sum)
	fmt.Println("平均分：", avg)
}
