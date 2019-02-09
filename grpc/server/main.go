package main

import (
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.Println("SERVER")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Could not create listener %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}
