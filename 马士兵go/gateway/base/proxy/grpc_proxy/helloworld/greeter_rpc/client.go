package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.用rpc链接服务器
	conn, err := rpc.Dial("tcp", "192.168.1.240:8004")
	if err != nil {
		fmt.Printf("dial err:%v\n", err)
	}
	defer conn.Close()

	// 2.调用远程函数
	var reply string
	if err = conn.Call("hello.HelloWorld", "派大星", &reply); err != nil {
		fmt.Println("call hello.HelloWorld err:", err)
		return
	}
	fmt.Println(reply)
}
