package main

// channel creation
var (
	firstChan chan interface{}
	// channel creation for reading
	secondChan <-chan interface{}
	// channel creation for writing
	thirdChan chan<- interface{}
)

func main() {
	// channel instantiation
	firstChan = make(chan interface{})

	// channel instantiation for reading
	secondChan = make(<-chan interface{})

	// channel instantiation for writing
	thirdChan = make(chan<- interface{})

	// this is valid because go implicitly converts channel from bidirectional to unidirectional
	thirdChan = firstChan
}
