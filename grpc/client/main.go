package main

import (
	"log"

	"../greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("CLIENT")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't open connection %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	log.Printf("Client: %f", c)
}
