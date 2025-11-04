package main

import (
	"fmt"
)

func maxOfThree(a int, b int, c int) int {
	if a > b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func main() {
	a := 1
	b := 2
	c := 3
	Three := maxOfThree(a, b, c)
	fmt.Println(Three)

}
