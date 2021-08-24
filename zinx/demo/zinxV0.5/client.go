package main

import (
	"fmt"
	"io"
	"net"
	"src/znet"
	"time"
)

/* 模拟客户端 */
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	// 1 连接远程服务器，获得conn
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// 发送封包msg消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessgaPackage(0, []byte("ZinxV0.5 client Test Message")))
		if err != nil {
			fmt.Println("Pack error:", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("Write error:", err)
			return
		}

		// 服务器回复msg消息，MsgId：1，ping..ping..ping

		// 1-先读取流中的head部分得到ID和datalen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}

		// 将二进制head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msghead error:", err)
			break
		}

		if msgHead.GetDataLen() > 0 {
			// 2-再根据datalen读取data
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error:", err)
				return
			}
			fmt.Println("---->recv msg id=", msg.Id, ", len=", msg.DataLen, ", data=", msg.Data)
		}

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
