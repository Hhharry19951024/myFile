package znet

import (
	"fmt"
	"net"
	"src/zinterface"
)

// IServer的接口实现，定义服务器模块
type Server struct {
	Name       string // 服务器名称
	IPVersrion string // 服务器IP版本
	IP         string // 服务器监听IP
	Port       int    // 服务器监听端口
	//当前Server添加一个router，server注册的链接对应处理业务
	Router zinterface.IRouter
}

func NewServer(name string) zinterface.IServer {
	s := &Server{
		Name:       name,
		IPVersrion: "tcp4",
		IP:         "0.0.0.0",
		Port:       8001,
		Router:     nil,
	}

	return s
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1获取tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersrion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}
		// 2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersrion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersrion, " err ", err)
			return
		}
		fmt.Println("start Zinx server succ, ", s.Name, "succ, Listening...")

		var cid uint32
		cid = 0

		// 3 阻塞连接，处理客户端业务
		for {
			// 如果有客户端连接，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				continue
			}

			// 将处理新链接的业务方法和conn进行绑定，得到链接模块
			dealConn := NewConnetcion(conn, cid, s.Router)
			cid++

			// 启动当前链接业务处理
			go dealConn.Start()

		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	// TODO
}

func (s *Server) Run() {
	// 启动server服务功能
	s.Start()

	// 可做额外业务

	// 阻塞状态
	select {}
}

func (s *Server) AddRouter(router zinterface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ")
}
