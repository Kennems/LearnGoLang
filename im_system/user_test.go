package main

import (
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func readConn(t *testing.T, conn net.Conn) string {
	t.Helper()

	_ = conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	buf, err := io.ReadAll(conn)
	if err != nil && !strings.Contains(err.Error(), "i/o timeout") {
		t.Fatalf("readConn() error = %v", err)
	}

	return string(buf)
}

func TestUserRename(t *testing.T) {
	server := NewServer("127.0.0.1", 8889)
	userConn, peerConn := net.Pipe()
	defer userConn.Close()
	defer peerConn.Close()

	user := NewUser(userConn, server)
	server.OnlineMap[user.Name] = user
	defer close(user.C)

	done := make(chan struct{})
	go func() {
		user.DoMessage("rename|alice")
		close(done)
	}()

	got := readConn(t, peerConn)
	<-done

	if user.Name != "alice" {
		t.Fatalf("user.Name = %q, want %q", user.Name, "alice")
	}
	if _, ok := server.OnlineMap["alice"]; !ok {
		t.Fatal("renamed user was not written back to OnlineMap")
	}
	if !strings.Contains(got, "您已更新用户名alice") {
		t.Fatalf("rename response = %q", got)
	}
}

func TestUserPrivateMessage(t *testing.T) {
	server := NewServer("127.0.0.1", 8889)

	senderConn, senderPeer := net.Pipe()
	receiverConn, receiverPeer := net.Pipe()
	defer senderConn.Close()
	defer senderPeer.Close()
	defer receiverConn.Close()
	defer receiverPeer.Close()

	sender := NewUser(senderConn, server)
	receiver := NewUser(receiverConn, server)
	defer close(sender.C)
	defer close(receiver.C)

	sender.Name = "alice"
	receiver.Name = "bob"
	server.OnlineMap[sender.Name] = sender
	server.OnlineMap[receiver.Name] = receiver

	done := make(chan struct{})
	go func() {
		sender.DoMessage("to|bob|hello")
		close(done)
	}()

	got := readConn(t, receiverPeer)
	<-done
	if !strings.Contains(got, "alice对你说：hello") {
		t.Fatalf("private message = %q", got)
	}
}
