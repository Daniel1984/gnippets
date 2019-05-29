package main

import (
	"./pipelines"
	"./stream"
	"fmt"
)

func main() {
	stream := stream.InitStream(&pipelines.PipelineOne{}, &pipelines.PipelineTwo{})

	go func() {
		stream.Receive(func(result int) {
			// the end result is received here after all pipeline transformations
			// have been applied
			fmt.Println(result)
			stream.Close()
		})
	}()

	// this is entry point where we send unprocessed data down the stream
	stream.Send(1)
	fmt.Scanln()
}
