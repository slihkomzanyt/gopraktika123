package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	if n < 0 {
		n = -n
	}
	str := strconv.Itoa(n)
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])

	}
	result, _ := strconv.Atoi(reversed)
	return result
}

func main() {

	fmt.Println(reverseNumber(123))
}
