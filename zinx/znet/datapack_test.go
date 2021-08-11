package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 负责测试datapack拆包封包的单元测试
func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	// 1-创建socketTCP
	listenner, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}

	// 创建一个go，承载负责从客户端处理业务
	go func() {
		// 2-从客户端读取数据，拆包处理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error: ", err)
			}

			go func(conn net.Conn) {
				// 处理客户端请求
				//-----------拆包---------
				// 定义拆包对象
				dp := NewDataPack()
				for {
					// 1，从conn读，读取包的head
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error: ", err)
						return
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error: ", err)
						return
					}

					if msgHead.GetDataLen() > 0 {
						// 有数据时需要第二次读取
						// 2，从coon读，根据head中的datalen读取data
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())

						// 根据DataLen从io流中读取
						_, err := io.ReadFull(conn, headData)
						if err != nil {
							fmt.Println("read head error: ", err)
							return
						}

						// 完整消息读取完毕
						fmt.Println("-->recv msgId= ", msg.Id, ", DataLen= ", msg.DataLen, ", Data= ", msg.Data)
					}

				}
			}(conn)
		}

	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("client dial error: ", err)
		return
	}
	// 创建一个封包对象dp
	dp := NewDataPack()

	// 模拟粘包过程，封装两个msg一同发送
	// 封装第一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack error: ", err)
		return
	}

	// 封装第二个msg2包
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack error: ", err)
		return
	}

	// 将两个包粘在一起
	sendData1 = append(sendData1, sendData2...)

	// 一次性发送服务端
	conn.Write(sendData1)

	// 客户端阻塞
	select {}
}
