package main

import (
	"context"
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

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Foo",
			LastName:  "Bar",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error makning request %v", err)
	}

	log.Printf("Got response and it's: %v", res.Result)
}
