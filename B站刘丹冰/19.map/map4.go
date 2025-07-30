package main

import "fmt"

// map是一个key(索引)value(值) 无序集合 主要查询方便
func main() {
	var coursesMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}
	//var coursesMap = map[string]string{} // nil map类型设值必须要初始化
	//var coursesMap1 = make(map[string]string, 3)
	fmt.Println(coursesMap["grpc"]) // 取值
	coursesMap["mysql"] = "mysql原理" // 放值
	fmt.Println(coursesMap)

	// 遍历
	// for _, value := range coursesMap {
	// 	fmt.Println(value)
	// }
	for key := range coursesMap {
		fmt.Println(key, coursesMap[key])
	}
	// map是无序的，而且不保证每次打印都是相同的顺序
	if d, ok := coursesMap["java"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("find", d)
	}

	fmt.Println(coursesMap["java"])

	// 删除元素
	delete(coursesMap, "grpc")
	fmt.Println(coursesMap)
	// map不是线程安全的
}
