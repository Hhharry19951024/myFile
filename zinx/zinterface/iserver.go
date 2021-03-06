package zinterface

// 定义服务器接口
type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 运行服务器
	Run()

	// 路由功能：给当前服务注册一个路由方法，供客户端链接处理使用
	AddRouter(router IRouter)
	
}
