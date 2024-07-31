package main

import (
	"os"
)

func main() {
	port := getPort()
	startServer(port)
}

func getPort() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return DEFAULT_PORT
}