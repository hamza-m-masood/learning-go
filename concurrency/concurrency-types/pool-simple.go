package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

// pool is only called twic
