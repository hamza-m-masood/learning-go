package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("Decrementing: %d\n", count)
	}

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("Incrementing: %d\n", count)
	}

	// increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	// decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}
