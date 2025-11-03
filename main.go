package main

import (
	"fmt"
)

func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false

}

func main() {
	year := 2024
	otvet := IsLeapYear(year)
	fmt.Println(otvet)

}
