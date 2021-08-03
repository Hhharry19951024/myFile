package znet

import (
	"fmt"
	"net"
	"src/zinterface"
)

type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn

	// 链接ID
	ConnID uint32

	// 当前链接状态
	isClosed bool

	// 当前链接绑定的API方法
	handleAPI zinterface.HandleFunc

	// 告知当前链接退出的channel
	ExitChan chan bool
}

// 初始化链接模块的方法
func NewConnetcion(conn *net.TCPConn, connID uint32, callbackAPI zinterface.HandlerFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 启动链接，准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID = ", c.ConnID)

	// 启动当前链接的读数据业务
	go c.StartReader()

	// TODO 启动当前链接的写数据业务

}

// 停止链接，结束链接工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 回收资源
	close(c.ExitChan)
}

// 获取当前链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据到客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//  读取客户端数据到buf,最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error ", err)
			continue
		}

		// 调用当前链接绑定的hanleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID = ", c.ConnID, " handle is error ", err)
			break
		}
	}
}
