package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip")
	flag.IntVar(&serverPort, "port", 8001, "set server port")
}

func NewClient(ServerIp string, ServerPort int) *Client {
	//create client
	client := &Client{
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		flag:       9,
	}

	//link server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ServerIp, client.ServerPort))
	if err != nil {
		fmt.Println("net Dial err:", err)
		return nil
	}
	client.conn = conn
	return client
}

func (client *Client) Menu() bool {
	fmt.Println("1-公聊模式")
	fmt.Println("2-私聊模式")
	fmt.Println("3-更改用户名")
	fmt.Println("0-退出")

	var flag int
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("===输入错误，请重试===")
		return false
	}
}

//handle server msg, diplay in Stdout
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) PublicChat() {
	fmt.Println("公聊模式中,exit退出...")
	var publicChatMsg string
	fmt.Scanln(&publicChatMsg)
	for publicChatMsg != "exit" {
		if len(publicChatMsg) != 0 {
			sendMsg := publicChatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}
		publicChatMsg = ""
		fmt.Println("公聊模式中,exit退出...")
		fmt.Scanln(&publicChatMsg)
	}
}

func (client *Client) SelectOnlineUser() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	client.SelectOnlineUser()

	fmt.Println("私聊模式中,请输入聊天用户名,exit退出...")
	var privateUserName string
	fmt.Scanln(&privateUserName)
	for privateUserName != "exit" {
		fmt.Println("私聊模式中,输入聊天信息,exit退出...")
		var privateChatMsg string
		fmt.Scanln(&privateChatMsg)
		for privateChatMsg != "exit" {
			if len(privateChatMsg) != 0 {
				sendMsg := "to|" + privateUserName + "|" + privateChatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}
			privateChatMsg = ""
			fmt.Println("私聊模式中,输入聊天信息,exit退出...")
			fmt.Scanln(&privateChatMsg)
		}
		client.SelectOnlineUser()
		fmt.Println("私聊模式中,请输入聊天用户名,exit退出...")
		fmt.Scanln(&privateUserName)
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名...")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.Menu() != true {
		}

		switch client.flag {
		case 1:
			client.PublicChat()
		case 2:
			client.PrivateChat()
		case 3:
			client.UpdateName()
		case 0:
			fmt.Println("退出")
		}
	}
}

func main() {
	flag.Parse() // command line parse

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>link server err")
		return
	}
	fmt.Println(">>>>>>>link server succ")

	//receive server msgResponse
	go client.DealResponse()

	//start client
	client.Run()
}
