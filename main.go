package main

import (
	"fmt"
)

func removeAtIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr

	} else {
		return append(arr[:index], arr[index+1:]...)
	}
}

func main() {
	nums := []int{1, 2, 3}
	nums = removeAtIndex(nums, 1)

	fmt.Println(nums)
}
