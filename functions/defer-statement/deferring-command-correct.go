package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	test()
}

func test() {
	time.Sleep(2 * time.Second)
}
