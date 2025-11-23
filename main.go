package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("тик")
		case <-ctx.Done():
			fmt.Println("stop")
			return

		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
}
