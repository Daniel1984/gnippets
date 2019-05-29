package pipelines

type PipelineTwo struct{}

func (p2 *PipelineTwo) Pipe(input chan int) chan int {
	output := make(chan int)

	go func() {
		for i := range input {
			// perform more transformations on the payload and send it further down
			// the stream
			output <- i * 2
		}

		close(output)
	}()

	return output
}
