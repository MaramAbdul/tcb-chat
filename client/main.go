package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const DEFAULT_PORT = "8989"

func main() {
	port := getPort()
	displayWelcomeMessage()
	connectToServer(port)
}

func getPort() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return DEFAULT_PORT
}

func displayWelcomeMessage() {
	welcomeMessage, err := ioutil.ReadFile("welcome.txt")
	if err != nil {
		log.Fatalf("Error reading welcome.txt: %v\n", err)
	}
	fmt.Print(string(welcomeMessage))
}

func connectToServer(port string) {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Could not connect to server: %v\n", err)
	}
	defer conn.Close()

	// fmt.Print("Enter your name: ")
	nameScanner := bufio.NewScanner(os.Stdin)
	nameScanner.Scan()
	name := nameScanner.Text()

	// Send the name to the server
	fmt.Fprintf(conn, "%s\n", name)

	go readMessages(conn)

	writeMessages(conn)
}

func readMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func writeMessages(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(conn, "%s\n", scanner.Text())
	}
}
