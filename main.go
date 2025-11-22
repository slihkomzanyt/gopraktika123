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

func kvadrat(nums <-chan int, out chan<- int) {

	for x := range nums {
		nums2 := x * x
		out <- nums2
	}
	close(out)
}

func printer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	nums := make(chan int)
	squared := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go generator(nums)
	go kvadrat(nums, squared)
	go printer(squared, &wg)

	wg.Wait()
}
