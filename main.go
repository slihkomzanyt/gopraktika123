package main

import (
	"fmt"
)

func sequenceGenerator() func() int {
	count := 0
	return func() int {
		count++
		return count
	}

}

func main() {
	next := sequenceGenerator()
	fmt.Println(next())
}
