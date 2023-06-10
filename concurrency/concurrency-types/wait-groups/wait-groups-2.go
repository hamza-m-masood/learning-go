package main

import (
	"fmt"
	"sync"
)

func main() {
	const numGreeters = 5
	var wg sync.WaitGroup

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
