package main

import (
	"fmt"
)

func reverse(s string) string {

	var result []rune
	originalRunes := []rune(s)
	for i := len(originalRunes) - 1; i >= 0; i-- {
		result = append(result, originalRunes[i])
	}
	return string(result)
}

func main() {

	fmt.Println(reverse("hello"))
	fmt.Println(reverse("привет"))
}
