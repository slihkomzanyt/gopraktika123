package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count int64
}

func (c *counter) inc() {

	atomic.AddInt64(&c.count, 1)

}

func (c *counter) value() int64 {

	return atomic.LoadInt64(&c.count)
}

func main() {
	c := counter{}

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
