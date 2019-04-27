package main

import (
	"./pipes"
	"./stream"
	"fmt"
)

func main() {
	stream := stream.InitStream(&pipes.PipeOne{})

	go func() {
		stream.Receive(func(result int) {
			fmt.Println(result)
		})
	}()

	stream.Send(1)
	fmt.Scanln()
}
