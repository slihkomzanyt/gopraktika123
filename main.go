package main

import (
	"fmt"
	"sort"
)

func mergeAndSort(arr1, arr2 []int) []int {
	rezultat := append(arr1, arr2...)
	sort.Ints(rezultat)
	return rezultat

}
func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 2, 6, 6, 1}
	rezultat := mergeAndSort(arr1, arr2)
	fmt.Println(rezultat)

}
