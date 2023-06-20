package main

import "fmt"

func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
}

// loopData is available outside of handleData but only handleData is using it.
// this is confinement through convention which is hard to keep because mistakes can be made.
