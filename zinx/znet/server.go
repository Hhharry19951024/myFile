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
}

func NewServer(name string) zinterface.IServer {
	s := &Server{
		Name:       name,
		IPVersrion: "tcp4",
		IP:         "0.0.0.0",
		Port:       8001,
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
		fmt.Println("start Zinx server succ, ", s.Name, "suucc, Listening...")

		// 3 阻塞连接，处理客户端业务
		for {
			// 如果有客户端连接，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				continue
			}

			// 客户端已建立连接，处理业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}

					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					// 结果回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						return
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	// TODO
}

func (s *Server) Run() {
	// 启动server服务功能
	s.Start()

	// 可做额外业务

	// 阻塞状态
	select {}
}
