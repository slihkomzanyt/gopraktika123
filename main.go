package main

import (
	"fmt"
	"math/rand"
)

func sumPart(nums []int, resultCh chan int) {
	part := 0
	for i := 0; i < len(nums); i++ {
		part += nums[i]
	}
	resultCh <- part
}

func main() {
	const parts = 10

	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	chunkSize := len(nums) / parts
	resultCh := make(chan int, parts)

	for i := 0; i < parts; i++ {
		start := i * chunkSize
		end := start + chunkSize
		chunk := nums[start:end]

		go sumPart(chunk, resultCh)
	}

	total := 0
	for i := 0; i < parts; i++ {
		val := <-resultCh
		total += val
	}

	fmt.Println("Total sum:", total)
}
