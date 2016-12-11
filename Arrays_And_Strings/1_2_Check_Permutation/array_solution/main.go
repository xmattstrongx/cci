package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPermutation("dog", "goD", true))
}

func IsPermutation(string1 string, string2 string, caseSensitive bool) bool {
	if !caseSensitive {
		string1 = strings.ToLower(string1)
		string2 = strings.ToLower(string2)
	}

	if len(string1) != len(string2) {
		return false
	}

	letters := make([]int, 128)
	for i := 0; i < len(string1); i++ {
		letters[string1[i]]++
	}

	for i := 0; i < len(string2); i++ {
		letters[string2[i]]--
		if letters[string2[i]] < 0 {
			return false
		}
	}

	return true
}
