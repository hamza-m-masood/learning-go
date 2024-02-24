package deferStatement

import (
	"fmt"
	"time"
)

func Run() {
	now := time.Now()
	defer fmt.Println(time.Since(now))
	test()
}

func test() {
	time.Sleep(2 * time.Second)
}
