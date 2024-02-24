package misc

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var mutex sync.Mutex
	counter := Counter{mutex, 12}
	counter.Increment()
	fmt.Println(counter.value)
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
