package main

import "fmt"

func main() {
	b := make(map[int]string)
	b[20095452] = "张三"
	b[20095387] = "李四"
	b[20098833] = "王五"
	fmt.Println(len(b))

	for k, v := range b {
		fmt.Println(k, v)
	}
	a := make(map[string]map[int]string)
	a["6班"] = make(map[int]string, 3)
	a["6班"][1] = "张三"
	a["6班"][2] = "李四"
	a["6班"][3] = "王五"
	a["6班"][4] = "赵六"
	fmt.Println(a["6班"])
	a["2班"] = make(map[int]string)
	a["2班"][1] = "小明"
	a["2班"][2] = "小花"
	a["2班"][3] = "小白"
	a["2班"][4] = "小红"
	fmt.Println(a["2班"])
	for k1, v1 := range a {
		for k2, v2 := range v1 {
			fmt.Printf("班级:%v 学生号为:%v 姓名:%v\n", k1, k2, v2)
		}
	}
}
