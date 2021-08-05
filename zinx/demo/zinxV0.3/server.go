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
func (p *PingRouter) PreHander(request zinterface.IRequest) {
	fmt.Println("Test Call Router PreHander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("Test Call Router PreHander err ", err)
	}
}

// Test Gandle
func (p *PingRouter) Hander(request zinterface.IRequest) {
	fmt.Println("Test Call Router Hander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("poing...ping...ping..."))
	if err != nil {
		fmt.Println("Test Call Router Hander err ", err)
	}
}

// Test PostHandle
func (p *PingRouter) PostHander(request zinterface.IRequest) {
	fmt.Println("Test Call Router PostHander")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
	if err != nil {
		fmt.Println("Test Call Router PostHander err ", err)
	}
}

func main() {
	// 创建server句柄，使用zinx的api
	s := znet.NewServer("[zinx V0.3]")

	// 给当强zinx框架添加自定义router
	s.AddRouter(&PingRouter{})

	//启动server
	s.Run()

}
