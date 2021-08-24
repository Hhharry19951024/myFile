package zinterface

/*
	IRequest接口：
	将客户端请求的链接信息， 与请求数据包装到一个Reques
*/

type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection

	// 得到请求消息数据
	GetData() []byte

	// 得到请求消息ID
	GetMsgID() uint32
}
