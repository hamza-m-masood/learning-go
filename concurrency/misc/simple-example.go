package misc

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	wg.Add(2)
	go sleep1()
	go sleep2()
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("elapsed time: %v\n", elapsed)
}

func sleep1() {
	defer wg.Done()
	time.Sleep(1 * time.Second)
}

func sleep2() {
	defer wg.Done()
	time.Sleep(1 * time.Second)
}
