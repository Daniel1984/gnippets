package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte(""))
	checkError(err)

	var buf [256]byte
	person := &Person{}
	n, err := conn.Read(buf[0:])
	checkError(err)
	json.Unmarshal(buf[0:n], person)
	fmt.Println("---->>>", person)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
