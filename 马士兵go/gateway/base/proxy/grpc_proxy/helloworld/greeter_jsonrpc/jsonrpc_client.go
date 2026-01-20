package main

import (
	"fmt"
	"gateway/base/proxy/grpc_proxy/helloworld/greeter_jsonrpc/inters"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyClient struct {
	c *rpc.Client
}

func (mc *MyClient) HelloWorld(arg string, reply *string) error {
	return mc.c.Call(inters.HelloServiceMethod, arg, &reply)
}

func NewClient(addr string) MyClient {
	//conn, err := rpc.Dial("tcp", addr)
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("dial err:%v\n", err)
	}
	return MyClient{c: conn}
}

func main() {
	// 1.用rpc链接服务器
	myClient := NewClient("192.168.1.240:8004")
	defer myClient.c.Close()

	// 2.调用远程函数
	var reply string
	err := myClient.HelloWorld("小白", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}
	fmt.Println(reply)
}
