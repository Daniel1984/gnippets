package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("CLIENT")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't open connection %v", err)
	}

	defer conn.Close()
}
