package main

import (
	"fmt"
	"net"
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
		// 2 连接调用Write写数据
		_, err := conn.Write([]byte("hello zinx V0.2..."))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server callback: %s, cnt = %d\n", buf, cnt)
		// cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
