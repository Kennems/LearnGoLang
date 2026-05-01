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

	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + " msg: " + msg
	this.Message <- sendMsg
	fmt.Printf(" 广播发送成功 %s \n", sendMsg)
}

func (this *Server) MessageListener() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("链接建立成功")
	//defer conn.Close()

	user := NewUser(conn, this)
	user.Online()

	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			// 去除 \n
			user.DoMessage(string(buf[:n-1]))

			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isLive:
			{

			}
		case <-time.After(time.Second * 60):
			{
				user.SendMsg("你被踢了")
				close(user.C)
				conn.Close()

				fmt.Printf("用户 %s 被踢了 \n", user.Name)
				// 退出当前handler
				return // runtime.Goexit()
			}
		}
	}
}

func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// close listener
	defer listener.Close()

	go this.MessageListener()

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		} else {
			fmt.Println("链接建立成功")
		}

		// do handler
		go this.Handler(conn)
	}
}
