package main

import (
	"fmt"
	"learn/rpc/new_hello_rpc/client_proxy"
)

func main() {
	//1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	// var reply *string = new(string)
	var reply string
	err := client.Hello("dawn", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
