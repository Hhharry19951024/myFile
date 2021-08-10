package znet

type Message struct {
	Id      uint32 // 消息ID
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

// 获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.Id
}

// 获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// 获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}

// 设置消息ID
func (m *Message) SetMsgID(id uint32) {
	m.Id = id
}

// 设置消息长度
func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

// 设置消息内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}
