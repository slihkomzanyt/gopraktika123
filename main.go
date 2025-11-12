package main

import (
	"fmt"
)

func mergetMap(m1, m2 map[string]int) map[string]int {
	mergetMap := make(map[string]int)
	for key, value := range m1 {
		mergetMap[key] = value
	}
	for key, value := range m2 {
		mergetMap[key] += value
	}
	return mergetMap
}

func main() {

	m1 := map[string]int{"apple": 5, "banana": 3}
	m2 := map[string]int{"banana": 5, "orenge": 3}
	result := mergetMap(m1, m2)
	fmt.Println(result)
}
