package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

//create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

// handle msg
func (pServer *Server) Handler(conn net.Conn) {
	//do something
	fmt.Println("server conn succ")
}

//start server interface
func (pServer *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", pServer.Ip, pServer.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// close listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		//do handler
		go pServer.Handler(conn)
	}
}
