package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    *sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	c := counter{
		mu: new(sync.Mutex),
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.inc()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.value())
}
