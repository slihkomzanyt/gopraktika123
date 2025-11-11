package main

import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	wordCount := make(map[string]int)
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

func main() {

	result := wordCount("Hello, hello world! World.")
	fmt.Println(result)

}
