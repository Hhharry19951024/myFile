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

	// 告知当前链接退出的channel
	ExitChan chan bool

	// 该链接处理的方法Router
	Router zinterface.IRouter
}

// 初始化链接模块的方法
func NewConnetcion(conn *net.TCPConn, connID uint32, router zinterface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 启动链接，准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID = ", c.ConnID)

	// 启动当前链接的读数据业务
	go c.StartReader()

	// TODO 启动当前链接的写数据业务
	for {
		select {
		case <-c.ExitChan:
			return
		}
	}
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

	// 通知缓冲队列读取数据业务，该链接已关闭
	c.ExitChan <- true

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

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//  读取客户端数据到buf,最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error ", err)
			c.ExitChan <- true
			continue
		}

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		// 执行注册的路由方法
		go func(request zinterface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// 从路由中找到注册绑定的Conn对应的router调用

	}
}
