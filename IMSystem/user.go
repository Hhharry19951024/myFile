package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	pServer *Server
}

// create a user api
func NewUser(conn net.Conn, pServer *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		C:       make(chan string),
		conn:    conn,
		pServer: pServer,
	}

	//run a user channel goroutine
	go user.ListenUserMsg()

	return user
}

func (user *User) Online() {
	//user onlinemap add user
	user.pServer.mapLock.Lock()
	user.pServer.OnlineMap[user.Name] = user
	user.pServer.mapLock.Unlock()

	//broadcast user online msg
	user.pServer.BroadcastMsg(user, "online")
}

func (user *User) Offline() {
	user.pServer.mapLock.Lock()
	delete(user.pServer.OnlineMap, user.Name)
	user.pServer.mapLock.Unlock()

	user.pServer.BroadcastMsg(user, "offline")
}

func (user *User) sendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

func (user *User) HandleMsg(msg string) {
	if msg == "who" {
		user.pServer.mapLock.Lock()
		for _, user := range user.pServer.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ": online...\n"
			user.sendMsg(onlineMsg)
		}
		user.pServer.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // format: rename| zhangsan
		newName := strings.Split(msg, "|")[1]

		_, ok := user.pServer.OnlineMap[newName]

		if ok {
			user.sendMsg("this userName has been used\n")
		} else {
			user.pServer.mapLock.Lock()
			delete(user.pServer.OnlineMap, user.Name)
			user.pServer.OnlineMap[newName] = user
			user.pServer.mapLock.Unlock()

			user.Name = newName
			user.sendMsg("you have renamed to " + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" { // format: to| zhangsan| msg
		acceptName := strings.Split(msg, "|")[1]
		acceptMsg := strings.Split(msg, "|")[2]
		if acceptName == "" {
			user.sendMsg("name fomat error, please use [to| zhangsan| msg]\n")
			return
		}
		if acceptMsg == "" {
			user.sendMsg("msg fomat error, please use [to| zhangsan| msg]\n")
			return
		}
		acceptUser, ok := user.pServer.OnlineMap[acceptName]
		if !ok {
			user.sendMsg("this user no exist\n")
			return
		}
		acceptUser.sendMsg("user[" + user.Name + "] to you: " + acceptMsg)

	} else {
		user.pServer.BroadcastMsg(user, msg)
	}
}

// listen user channel
func (user *User) ListenUserMsg() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}
