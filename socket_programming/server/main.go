package main

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = 8000
	TYPE = "tcp"
)

func main() {
	uri := fmt.Sprintf("%v:%v", HOST, PORT)
	fmt.Printf("starting server at %v", uri)
	server, err := net.Listen(TYPE, uri)
	if err != nil {
		log.Fatal("Error starting server:", err.Error())
	}

	defer server.Close()
	fmt.Println("waiting for connections")

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("can't handle connection:", err.Error())
		}

		fmt.Println("Handling connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buff := make([]byte, 1024)
	msgLen, err := conn.Read(buff)
	if err != nil {
		log.Fatal("can read msg:", err.Error())
	}
	fmt.Println("Received:", string(buff[:msgLen]))
	conn.Close()
}
