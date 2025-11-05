package main

import (
	"fmt"
	"unicode"
)

func charCase(r rune) string {

	switch {
	case unicode.Is(unicode.Latin, r) && unicode.IsUpper(r):
		return ("Латинская заглавная")

	case unicode.Is(unicode.Cyrillic, r) && unicode.IsUpper(r):
		return ("Кириллица заглавная")

	case unicode.Is(unicode.Latin, r) && unicode.IsLower(r):
		return ("латинская маленькая")

	case unicode.Is(unicode.Cyrillic, r) && unicode.IsLower(r):
		return ("кириллица маленькая")

	default:
		return ("Другое")
	}

}

func main() {

	fmt.Println(charCase('h'))
	fmt.Println(charCase('G'))
}
