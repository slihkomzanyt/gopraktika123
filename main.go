package main

import (
	"fmt"
)

func increment(ptr *int) {

	*ptr++

}

func main() {
	n := 4
	increment(&n)
	fmt.Println(n)
}
