package server_proxy

import (
	"learn/rpc/new_hello_rpc/handler"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv *handler.NewHelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
