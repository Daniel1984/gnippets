package pipes

type PipeTwo struct{}

func (p2 *PipeTwo) Pipe(input chan int) chan int {
	output := make(chan int)

	go func() {
		for i := range input {
			output <- i * 2
		}

		close(output)
	}()

	return output
}
