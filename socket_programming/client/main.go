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
	fmt.Printf("Client is running at: %v\n", uri)
	conn, err := net.Dial(TYPE, uri)

	if err != nil {
		log.Fatal("Error dialing:", err.Error())
	}

	fmt.Println("sending data")
	_, err = conn.Write([]byte("Hello world!"))

	if err != nil {
		log.Fatal("Error sending message:", err.Error())
	}

	defer conn.Close()
}
