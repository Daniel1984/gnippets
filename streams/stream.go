package streams

type pipe struct {
	input  chan int
	output chan int
}

func (p *pipe) Send(e int) {
	p.input <- e
}

func (p *pipe) Receive(callback func(int)) {
	for e := range p.output {
		callback(e)
	}
}

func (p *pipe) Close() {
	close(p.input)
}

type Stream interface {
	Pipe(input chan int) chan int
}

func InitStream(filters ...Stream) *pipe {
	source := make(chan int)
	var nextStream chan int

	for _, filter := range filters {
		if nextStream == nil {
			nextStream = filter.Pipe(source)
		} else {
			nextStream = filter.Pipe(nextStream)
		}
	}

	return &pipe{input: source, output: nextStream}
}
