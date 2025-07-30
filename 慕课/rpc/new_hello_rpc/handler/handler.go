package handler

const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request // 返回值是通过reply的值
	return nil
}
