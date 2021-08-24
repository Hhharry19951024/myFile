package main

import (
	"fmt"
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

		// 先读取流中的head部分得到ID和datalen

		// 再根据datalen读取data

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
