package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, d time.Duration, msg string, ch chan<- string) {
	time.Sleep(d)

	select {
	case <-ctx.Done():

		return
	case ch <- msg:

	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan string)

	go worker(ctx, time.Second, "a", ch)
	go worker(ctx, 3*time.Second, "b", ch)

	res := <-ch
	fmt.Println(res)

	cancel()
}
