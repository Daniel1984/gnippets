package stream

type Stream struct {
	input  chan int
	output chan int
}

func (s *Stream) Send(e int) {
	s.input <- e
}

func (s *Stream) Receive(callback func(int)) {
	for transformedPayload := range s.output {
		callback(transformedPayload)
	}
}

func (s *Stream) Close() {
	close(s.input)
}

// Pipeliner interface has only one Pipe method in this case it operates
// over channel that accepts int type of value (could be anything)
type Pipeliner interface {
	Pipe(input chan int) chan int
}

func InitStream(pipelines ...Pipeliner) *Stream {
	inputChan := make(chan int)
	var outputChan chan int

	// first pipeline will get initial inputChannel as input for data to flow
	// each pipelines Pipe method returns output channel that becomes an input
	// for next pipeline in queue which guarantees sequential execution. All
	// depends in what order you pass in pilelines to InitStream function
	for _, p := range pipelines {
		if outputChan == nil {
			outputChan = p.Pipe(inputChan)
		} else {
			outputChan = p.Pipe(outputChan)
		}
	}

	return &Stream{input: inputChan, output: outputChan}
}
