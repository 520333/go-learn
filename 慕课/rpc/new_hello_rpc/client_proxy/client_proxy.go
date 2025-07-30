package client_proxy

import (
	"learn/rpc/new_hello_rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("连接失败")
	}
	return HelloServiceStub{conn}
}

// request，reply
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
