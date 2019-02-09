package main

import (
	"google.golang.org/grpc"
	"log"
)

type server struct{}

func main() {
	log.Println("SERVAER")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Could not create listener %v", err)
	}

	s := grpc.NewServer()
}
