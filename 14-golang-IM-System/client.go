package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

var reader *bufio.Reader

func NewClient(serverIp string, serverPort int) *Client {

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       9999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	return client
}

func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1、公聊模式")
	fmt.Println("2、私聊模式")
	fmt.Println("3、更新用户名")
	fmt.Println("0、退出")
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("输入有误，请重新输入")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名:")
	client.Name, _ = reader.ReadString('\n')
	client.Name = strings.TrimSpace(client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("UpdateName error:", err)
		return false
	}

	return true
}

func (client *Client) PublicChat() {
	var chatMsg string
	fmt.Println("请输入聊天内容，exit退出:")
	chatMsg, _ = reader.ReadString('\n')
	chatMsg = strings.TrimSpace(chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("write error:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println("请输入聊天内容，exit退出:")
		chatMsg, _ = reader.ReadString('\n')
		chatMsg = strings.TrimSpace(chatMsg)
	}
}

func (client *Client) selectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("selectUsers error:", err)
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.selectUsers()
	fmt.Println("请输入聊天对象的用户名， exit退出")
	remoteName, _ = reader.ReadString('\n')
	remoteName = strings.TrimSpace(remoteName)

	for remoteName != "exit" {
		fmt.Println("请输入聊天内容，exit退出:")
		chatMsg, _ = reader.ReadString('\n')
		chatMsg = strings.TrimSpace(chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("write error:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println("请输入聊天内容，exit退出:")
			chatMsg, _ = reader.ReadString('\n')
			chatMsg = strings.TrimSpace(chatMsg)
		}

		client.selectUsers()
		fmt.Println("请输入聊天对象，exit退出")
		remoteName, _ = reader.ReadString('\n')
		remoteName = strings.TrimSpace(remoteName)
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式")
			client.PrivateChat()
			break
		case 3:
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip")
	flag.IntVar(&serverPort, "port", 8889, "set server port")
}

func main() {
	// 命令行解析
	flag.Parse()
	reader = bufio.NewReader(os.Stdin) // 初始化全局 reader

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("Connect to the server successfully")
		return
	}

	go client.DealResponse()

	fmt.Println("Connect to the server success")
	client.Run()
}
