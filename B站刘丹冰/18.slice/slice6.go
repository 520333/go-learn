package main

/*
import "fmt"

// slice是数组的弱化 实现也是基于数组
// slice初始化3种方式: 1.从数组创建 2.使用string{} 3.使用make
func main() {
	// 方法1
	var courses []string
	fmt.Printf("类型:%T, 长度:%v\r\n", courses, len(courses))
	courses = append(courses, "go")
	courses = append(courses, "java")
	courses = append(courses, "python")
	fmt.Printf("类型:%T, 长度:%v,数据:%s\r\n", courses, len(courses), courses)
	fmt.Println(courses[2])

	// 方法2 string{}
	test := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	a1 := test[0:len(test)] //左闭右开 python语法
	fmt.Println(a1)

	// 方法3 make
	skill := make([]string, 3)
	skill[0] = "go"
	fmt.Println(skill)

	skill1 := make([]string, 3) // 如果声明了长度 只能使用append方法追加 不可用使用[index]方式
	skill1[0] = "go"
	skill1[1] = "grpc"
	skill1[2] = "gin"
	skill1 = append(skill1, "python")
	fmt.Println(skill1)

	// 访问切片元素
	fmt.Println(skill1[1])
	fmt.Println(skill1[1:]) //[start:end] 如果只有start没有end 表示从start开始到结尾的所有数据
	fmt.Println(skill1[:3]) //[start:end] 如果只有end没有start 表示从0到end的之前的所有数据
	fmt.Println(skill1[:])  //[start:end] 所有数据

	language := []string{"go", "grpc"}
	language2 := []string{"mysql", "es", "gin"}
	//language = append(language, "php", "java", "grovvy")
	for _, value := range language2 {
		language = append(language, value)
	}
	language = append(language, language2[1:]...)
	fmt.Println(language)

	language3 := []string{"go", "grpc", "mysql", "es", "gin"}
	myLanguage := append(language3[:2], language3[3:]...)
	fmt.Println(myLanguage)
	language3 = language3[:3]
	fmt.Println(language3)

	// 复制slice
	copyLanguage1 := language3
	copyLanguage2 := language3[:]
	fmt.Println(copyLanguage1, copyLanguage2)

	var language4 = make([]string, len(language3))
	copy(language4, language3)

	fmt.Println("===============")
	language3[0] = "java"
	fmt.Println(language4)
	fmt.Println(copyLanguage2)

}
*/
