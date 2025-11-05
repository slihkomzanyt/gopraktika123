package main

import (
	"fmt"
)

func calculator(a, b int, op string) int {

	errorMsg := "на ноль делить нельзя"
	switch op {
	case "+":
		return a + b

	case "-":
		return a - b

	case "*":
		return a * b

	case "/":
		if b == 0 {
			fmt.Println(errorMsg)
			return 0
		} else {
			return a / b
		}

	default:
		return 0
	}

}

func main() {
	a := 1
	b := 2

	fmt.Println(calculator(a, b, "+"))

}
