package main

// myfunc updateString(s string) {
// 	s = "这是一个新值"
// }
// myfunc updateStringWithPointer(s *string) {
// 	*s = "这是一个string类型的指针新值"
// }

// myfunc main() {
// 	var s string
// 	s = "这是一个字符串"
// 	fmt.Println("变量s的内存地址是:", &s) // stdout: 变量s的内存地址是: 0xc000088240

// 	sPoint := &s
// 	fmt.Println("指针sp:", sPoint) //stdout: 指针sp: 0xc000088240

// 	var sp2 *string               //声明一个string类型的指针
// 	fmt.Println("指针sp2未赋值:", sp2) //stdout: 指针sp2未赋值: <nil>

// 	sp2 = &s
// 	fmt.Println("指针sp2:", sp2) //stdout: 指针sp2: 0xc000088240

// 	fmt.Println("指针对应内存地址的值:", *sp2) //stdout: 指针对应内存地址的值: 这是一个字符串

// 	updateString(s)
// 	fmt.Println("修改后的s:", s) //stdout: 修改后的s: 这是一个字符串

// 	updateStringWithPointer(&s)
// 	fmt.Println("真正修改后的的s:", s) //stdout: 真正修改后的的s: 这是一个string类型的指针新值
// }
