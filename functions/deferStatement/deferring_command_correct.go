package deferStatement

import (
	"fmt"
	"time"
)

func RunCorrect() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	testCorrect()
}

func testCorrect() {
	time.Sleep(2 * time.Second)
}
