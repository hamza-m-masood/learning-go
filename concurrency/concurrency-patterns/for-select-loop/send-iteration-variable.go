package main

import "fmt"

func main() {
	done := make(chan int)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream := <-s:
			fmt.Println("stringstream", stringStream)
		}
	}
}
