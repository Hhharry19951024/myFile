package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//online user list
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//broadcast chan
	Msg chan string
}

//create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Msg:       make(chan string),
	}
	return server
}

//broadcast msg
func (pServer *Server) BroadcastMsg(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg
	pServer.Msg <- sendMsg
}

func (pServer *Server) ListenMsg() {
	for {
		msg := <-pServer.Msg

		// send msg to online users
		pServer.mapLock.Lock()
		for _, cli := range pServer.OnlineMap {
			cli.C <- msg
		}
		pServer.mapLock.Unlock()
	}
}

// handle msg
func (pServer *Server) Handler(conn net.Conn) {
	fmt.Println("server conn succ")

	user := NewUser(conn, pServer)

	user.Online()

	//create user active channel
	isLive := make(chan bool)

	//read client msg, broadcast msg to online users
	go func() {
		buf := make([]byte, 4096)
		for {
			readLength, err := conn.Read(buf)
			if readLength == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			// handle "\n"
			msg := string(buf[:readLength-1])
			user.HandleMsg(msg)

			isLive <- true
		}
	}()

	//handler block
	for {
		select {
		case <-isLive:
			// do nothing
		case <-time.After(time.Second * 1000):
			user.sendMsg("you are deleted")

			close(user.C) // close channel
			conn.Close()  // close conn
			return
		}
	}
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

	// do listen msg
	go pServer.ListenMsg()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		// do handler
		go pServer.Handler(conn)
	}
}
