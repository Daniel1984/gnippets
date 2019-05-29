package pipelines

type PipelineOne struct{}

func (p1 *PipelineOne) Pipe(input chan int) chan int {
	output := make(chan int)

	go func() {
		for i := range input {
			// make whatever modifications to payload, which in this case is i:int
			// and send it down the stream to next pipeline
			i = i + 2
			output <- i
		}

		close(output)
	}()

	return output
}
