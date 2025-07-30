package main

/*
空接口：空接口不会定义任何的方法，所以，无论什么类型都实现了这个接口
空接口是可以接收任何类型的参数值
*/
// type EmptyInterface interface{}

// func main() {
// 	fmt.Println("打印一条数据")
// 	var ei EmptyInterface
// 	s1 := "这是一串字符串"
// 	ei = s1
// 	fmt.Println("ei:", ei)

// 	//定义空接口方法2
// 	var ei2 interface{}
// 	ei2 = s1
// 	fmt.Println("ei2:", ei2)
// 	//map 联系方式:手机号(int),座机号(010-12345678)
// 	contacts := make(map[string]interface{})
// 	contacts["chuang"] = "180-1122-5566"
// 	contacts["obiwang"] = "010-12345678"
// 	fmt.Println(contacts)
// }
