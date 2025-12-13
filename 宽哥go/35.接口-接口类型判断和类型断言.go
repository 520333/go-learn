package main

// myfunc dealData(data interface{}) {
// 	t, ok := data.(string) //如果data是string类型的，ok=true t=data
// 	if ok {
// 		fmt.Println("当前类型为string,变量的值是:", t)
// 	} else {
// 		fmt.Println("data不是字符串")
// 		fmt.Println("当前t的值:", t)
// 	}
// }
// myfunc getType(i interface{}) {
// 	switch t := i.(type) {
// 	case int:
// 		fmt.Println("当前值为int类型:", t)
// 	case string:
// 		fmt.Println("当前值为int类型:", t)
// 	case float32:
// 		fmt.Println("当前值为float类型:", t)
// 	case bool:
// 		fmt.Println("当前值为bool类型:", t)
// 	default:
// 		fmt.Println("当前类型不在处理范围")
// 	}

// }
// myfunc main() {
// 	//类型断言：大致知道了接口可能是某种类型，然后使用t,ok:=i.(string)
// 	s := "hello"
// 	dealData(s) //stdout: 当前类型为string,变量的值是: hello
// 	i := 123456
// 	dealData(i)   //stdout: data不是字符串
// 	getType(i)    //stdout: 当前值为int类型: 123456
// 	getType(true) //stdout: 当前值为bool类型: true
// 	var a interface{}
// 	getType(a) //stdout: 当前类型不在处理范围:

// }
