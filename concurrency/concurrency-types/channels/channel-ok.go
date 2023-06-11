package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	salutations, ok := <-stringStream
	fmt.Println(salutations, ok)
}
