package inters

// HelloServiceName 服务名
const HelloServiceName = "HELLO"

// HelloServiceMethod 服务方法
const HelloServiceMethod = "HELLO.HelloWorld"

// MyInterface 定义服务方法
type MyInterface interface {
	// HelloWorld 定义服务函数
	// arg: 传入参数
	// reply: 传出参数 指针类型
	HelloWorld(arg string, reply *string) error
}
