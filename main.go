package main

import (
	"fmt"
	"sync"
)

func generator(nums chan<- int) {
	for x := 1; x <= 5; x++ {
		nums <- x
	}
	close(nums)
}

func consumer(nums <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range nums {
		fmt.Println(num)
	}
}

func main() {
	nums := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go generator(nums)
	go consumer(nums, &wg)

	wg.Wait()
}
