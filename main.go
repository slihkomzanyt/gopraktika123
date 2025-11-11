package main

import (
	"fmt"
)

func safeFunction(f func()) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("ошибка: %v\n", r)
		}

	}()

	f()

}

func main() {

	safeFunction(func() {
		fmt.Println("code go")
		panic("code trabl")
	})
}
