package main

import (
    "net"
)

func broadcast(message string, sender net.Conn) {
    mu.Lock()
    chatHistory = append(chatHistory, message)
    for client := range clients {
        if client != sender {
            client.Write([]byte(message))
        }
    }
    mu.Unlock()
}

func sendChatHistory(conn net.Conn) {
    mu.Lock()
    for _, message := range chatHistory {
        conn.Write([]byte(message))
    }
    mu.Unlock()
}