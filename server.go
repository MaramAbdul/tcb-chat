package main

import (
	"log"
	"net"
	"sync"
)

const (
	DEFAULT_PORT = "8989"
	MAX_CLIENTS  = 10
)

var (
	clients     = make(map[net.Conn]string)
	chatHistory []string
	mu          sync.Mutex
)

func startServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[USAGE]: ./TCPChat $port\n")
	}
	defer ln.Close()
	log.Printf("Server started on port %s\n", port)

	for {
		if len(clients) < MAX_CLIENTS {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConnection(conn)
		}
	}
}
