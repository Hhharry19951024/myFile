package main

import (
	"fmt"
	"src/zinterface"
	"src/znet"
)

/*
	基于Zinx框架开放 服务器应用程序
*/

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHander
func (p *PingRouter) PreHandle(request zinterface.IRequest) {
	fmt.Println("Test Call Router PreHander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("Test Call Router PreHander err ", err)
	}
}

// Test Gandle
func (p *PingRouter) Handle(request zinterface.IRequest) {
	fmt.Println("Test Call Router Hander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("poing...ping...ping...\n"))
	if err != nil {
		fmt.Println("Test Call Router Hander err ", err)
	}
}

// Test PostHandle
func (p *PingRouter) PostHandle(request zinterface.IRequest) {
	fmt.Println("Test Call Router PostHander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("Test Call Router PostHander err ", err)
	}
}

func main() {
	// 创建server句柄，使用zinx的api
	s := znet.NewServer("[zinx V0.4]")

	// 给当强zinx框架添加自定义router
	s.AddRouter(&PingRouter{})

	//启动server
	s.Run()

}
