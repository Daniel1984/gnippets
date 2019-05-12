package pipes

type PipeOne struct{}

func (p1 *PipeOne) Pipe(input chan int) chan int {
	output := make(chan int)

	go func() {
		for i := range input {
			output <- i
			output <- i
			output <- i
			output <- i
		}

		close(output)
	}()

	return output
}
