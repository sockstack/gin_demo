package hello

import "fmt"

//原服务
type HelloService struct {
}

func (h *HelloService)SayHello(username string) string {
	return fmt.Sprintf("hello, %s", username)
}

//新服务
type GreeterService struct {
}

func (h *GreeterService)SayHello(username string) string {
	return fmt.Sprintf("hello, %s, how are you!", username)
}
