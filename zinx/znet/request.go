package znet

import "src/zinterface"

type Request struct {
	// 已经和客户端建立链接
	conn zinterface.IConnection

	// 客户端请求数据
	msg zinterface.IMessage
}

// 得到当前链接
func (r *Request) GetConnection() zinterface.IConnection {
	return r.conn
}

// 得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
