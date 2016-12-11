package main

import (
	"fmt"
)

func main() {
	word := "hunkydor"
	fmt.Println(IsUnique(word))
}

func IsUnique(word string) bool {
	if len(word) > 128 {
		return false
	}

	charMap := make(map[rune]int)
	for _, letter := range word {
		_, ok := charMap[letter]
		if ok {
			return false
		}
		charMap[letter] = 1

	}

	return true
}
