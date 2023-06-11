package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
		stringStream <- "what are you doing?"
		close(stringStream)
	}()
	for i := range stringStream {
		fmt.Println(i)
	}
}
