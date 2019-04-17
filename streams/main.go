package main

import (
	"./pipes"
	"./stream"
	"fmt"
)

func main() {
	pipe1 := pipelines.NewPipe1()
	pipe2 := pipelines.NewPipe2()

	stream := stream.InitStream(pipe1, pipe2)

	go func() {
		stream.Receive(func(result int) {
			fmt.Println(result)
		})
	}()

}
