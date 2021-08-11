package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"src/utils"
	"src/zinterface"
)

// 封包拆包的具体实现
type DataPack struct{}

// 封包拆包实例的初始化
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的头的长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// Datelen uint32(4字节) + ID uint32(4字节)
	return 8
}

// 封包方法
func (dp *DataPack) Pack(msg zinterface.IMessage) ([]byte, error) {
	// 创建一个存放bytes的缓冲
	dataBuf := bytes.NewBuffer([]byte{})

	// 将DataLen写进dataBuf中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 将Id写进dataBuf中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 将Data写进dataBuf中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

// 拆包方法(将head信息读取DataLen)，然后进行拆包
func (dp *DataPack) Unpack(binaryData []byte) (zinterface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuf := bytes.NewReader(binaryData)

	// 解压head信息，得到DataLen，Id
	msg := &Message{}

	// 读取DataLen
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读取Id
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断DataLen是否合法
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("msg recv datalen too large!!!")
	}

	return msg, nil
}
