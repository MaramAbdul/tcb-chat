package main

import (
    "bufio"
    "net"
)

func getInput(conn net.Conn) string {
    scanner := bufio.NewScanner(conn)
    if scanner.Scan() {
        return scanner.Text()
    }
    return ""
}

func getInputWithErr(conn net.Conn) (string, error) {
    scanner := bufio.NewScanner(conn)
    if scanner.Scan() {
        return scanner.Text(), nil
    }
    return "", scanner.Err()
}