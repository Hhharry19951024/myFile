package main

import (
	"zinx/znet"
)

func main() {
	// 创建server句柄，使用zinx的api
	s := znet.NewServer("[zinx V0.1]")
	//启动server
	s.Run()

}
