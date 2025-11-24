package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	mu    sync.RWMutex
	value int
}

func (s *SafeNumber) Get() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *SafeNumber) Set(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value = v
}

func main() {
	num := &SafeNumber{value: 0}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readers := 5
	for i := 1; i <= readers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("reader %d: stop\n", id)
					return
				default:
					v := num.Get()
					fmt.Printf("reader %d: value = %d\n", id, v)
					time.Sleep(200 * time.Millisecond)
				}
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			num.Set(i)
			fmt.Printf("writer: set value = %d\n", i)
		}
		cancel()
	}()

	wg.Wait()
}
