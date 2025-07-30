package main

import (
	"fmt"
)

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	company Company
}
type PrintResult struct {
	Info string
	Err  error
}

func rpcPrintln(employee Employee) PrintResult {
	//1.tcp/http建立连接【客户端】
	//2.将employee对象序列化成json字符串【服务端】
	//3.发送json字符串【服务端】
	//4.将服务端返回的数据反序列化json字符串【客户端】
	return PrintResult{}
}

func main() {
	fmt.Println(Add(5, 10))
	fmt.Println(Employee{
		Name: "dawn",
		company: Company{
			Name:    "dawn",
			Address: "福建省",
		},
	})
}
