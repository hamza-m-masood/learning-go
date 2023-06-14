package main

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select {
	case <-c1:
	// do something
	case <-c2:
	// do something
	case <-c3:
		// do something
	}
}
