package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	name := getInput(conn)
	if name == "" {
		return
	}

	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	broadcast(fmt.Sprintf("%s has joined the chat\n", name), conn)
	sendChatHistory(conn)

	go readMessages(conn, name)

	for {
		time.Sleep(1 * time.Second)
	}
}

func readMessages(conn net.Conn, name string) {
	for {
		message, err := getInputWithErr(conn)
		if err != nil {
			handleDisconnection(conn, name)
			break
		}
		if message == "" {
			continue
		}
		broadcast(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, message), conn)
	}
}

func handleDisconnection(conn net.Conn, name string) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	broadcast(fmt.Sprintf("%s has left the chat\n", name), nil)
}
