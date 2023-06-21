package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}

// This is an example of confinement using a data structure that is not concurrent safe
// a concurrent safe datastructure would be a channel
// printData is not closing over the slice, so it is not using the same data. It is using it's own copy.
