package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t1 Teacher // 创建方式1
	t1.Name = "海绵宝宝"
	t1.Age = 20
	t1.School = "海绵宝宝综合大学"
	fmt.Println(t1)

	var t2 *Teacher = new(Teacher) // 创建方式2
	(*t2).Name = "派大星"
	(*t2).Age = 21
	t2.School = "派大星艺术学院"
	fmt.Println(*t2)

	var t3 Teacher = Teacher{Name: "蟹老板", Age: 30, School: "海绵宝宝综合大学"} // 创建方式3
	fmt.Println(t3)

	var t4 *Teacher = &Teacher{"章鱼哥", 21, "八加九理工大学"} // 创建方式3
	fmt.Println(*t4)

}
