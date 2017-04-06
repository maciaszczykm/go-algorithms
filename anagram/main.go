package main

import (
	"fmt"
)

func main() {
	check("not", "anagram")
	check("test", "tset")
	check("joke", "ekoi")
	check("joke", "ekoj")
}

// check performs check if two provided words are anagrams and displays result.
func check(first, second string) {
	if isAnagram(first, second) {
		fmt.Printf("%s is an anagram of %s\n", first, second)
	} else {
		fmt.Printf("%s is not an anagram of %s\n", first, second)
	}
}

// isAnagram returns true if two provided words are anagrams.
func isAnagram(first, second string) bool {
	length := len(first)
	if length != len(second) {
		return false
	}

	for i := 0; i < length; i++ {
		if first[i] != second[length - i - 1] {
			return false
		}
	}

	return true
}