package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(IsPermutation("dog", "goD", false))
}

func IsPermutation(string1 string, string2 string, caseSensitive bool) bool {
	if !caseSensitive {
		string1 = strings.ToLower(string1)
		string2 = strings.ToLower(string2)
	}

	if len(string1) != len(string2) {
		return false
	}

	charMap1 := make(map[rune]int)
	for _, letter := range string1 {
		_, ok := charMap1[letter]
		if !ok {
			charMap1[letter] = 1
		} else {
			charMap1[letter]++
		}
	}

	charMap2 := make(map[rune]int)
	for _, letter := range string2 {
		_, ok := charMap2[letter]
		if !ok {
			charMap2[letter] = 1
		} else {
			charMap2[letter]++
		}
	}

	eq := reflect.DeepEqual(charMap1, charMap2)
	if !eq {
		return false
	}
	return true
}
