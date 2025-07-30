package main

// func main() {
// 	teacherAge := make(map[string]int)
// 	fmt.Println("对象的初始化值:", teacherAge) //stdout: 对象的初始化值: map[]

// 	teacherAge["dawn"] = 18
// 	teacherAge["chuang"] = 20
// 	teacherAge["obiwan"] = 21
// 	fmt.Println("赋值后的值:", teacherAge) //stdout: 赋值后的值: map[chuang:20 dawn:18 obiwan:21]

// 	teacherAge2 := map[string]int{ //在声明变量的时候直接进行赋值操作
// 		"age":    18,
// 		"salary": 10240,
// 	}
// 	fmt.Println("teacherAge2的值:", teacherAge2) //stdout: teacherAge2的值: map[age:18 salary:10240]

// 	var teacherAddress map[string]string //使用var赋值
// 	teacherAddress = make(map[string]string)
// 	teacherAddress["dawn"] = "北京"
// 	teacherAddress["chuang"] = "福建"
// 	fmt.Println("teacherAddress的值:", teacherAddress) //stdout: teacherAddress的值: map[chuang:福建 dawn:北京]

// 	// 增删改查--查
// 	fmt.Println("查询dawn老师的地址:", teacherAddress["dawn"]) //stdout: 查询dawn老师的地址: 北京
// 	searchName := "chuang"
// 	fmt.Printf("查询老师:%s的地址:%s", searchName, teacherAddress[searchName]) //stdout: 查询老师:chuang的地址:福建
// 	for k, v := range teacherAge {
// 		fmt.Printf("老师:%s,年龄:%d", k, v) //stdout: 老师:chuang,年龄:20老师:obiwan,年龄:21
// 	}
// 	fmt.Println("取一个不存在的值:", teacherAddress["eeeee"]) //stdout: 取一个不存在的值:
// 	address, ok := teacherAddress["dawn"]
// 	if ok {
// 		fmt.Println("能查询到", ok, address) //stdout: 能查询到 true 北京
// 	} else {
// 		fmt.Println("找不到", ok, address)
// 	}

// 	// 增删改查--改
// 	teacherAddress["dawn"] = "上海"
// 	fmt.Println("修改后的值:", teacherAddress["dawn"]) //stdout: 修改后的值: 上海

// 	// 增删改查--删
// 	delete(teacherAddress, "dawn")
// 	fmt.Println("删除后的值:", teacherAddress) //stdout: 删除后的值: map[chuang:福建]

// }
