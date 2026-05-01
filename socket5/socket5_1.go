package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Errorf("accept error: %w", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		buf, err := reader.ReadByte()
		if err != nil {
			fmt.Errorf("read buf error: %w", err)
			return
		}
		_, err = conn.Write([]byte{buf})
		if err != nil {
			fmt.Errorf("write buf back error: %w", err)
			return
		}
	}
}
