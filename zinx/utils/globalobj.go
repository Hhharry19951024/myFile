package utils

import (
	"encoding/json"
	"io/ioutil"
	"src/zinterface"
)

/*
	存储有关zinx框架的全局参数，供其他模块使用
	一些参数可由zinx.json配置
*/
type GlobalObj struct {
	/*
		Server
	*/
	TcpServer zinterface.IServer // 当前zinx全局的server对象
	Host      string             // 当前服务器主机监听的IP
	TcpPort   int                // 当前服务器主机监听的端口号
	Name      string             // 当前服务器的名称

	/*
		Zinx
	*/
	Version        string // 当前zinx的版本号
	MaxConn        int    // 当前服务器主机允许的最大链接数
	MaxPackageSize uint32 // 当前zinx框架数据包的最大值
}

/*
	定义一个全局的对外Globalobj
*/
var GlobalObject *GlobalObj

/*
	从zinx.json加载自定义参数
*/
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将json文件解析到struct
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	提供一个init方法，初始化当前的GlobalObject
*/
func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerAPP",
		Version:        "V0.4",
		TcpPort:        8001,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 尝试conf/zinx.json加载自定义参数
	// GlobalObject.Reload()
}
