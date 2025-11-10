package main

import (
	"fmt"
)

func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            result = 0  
        }
    } 
    
    if b == 0 {
        panic("деление на ноль")
    }
    return a / b
}

func main() {
	
	fmt.Println()
}
