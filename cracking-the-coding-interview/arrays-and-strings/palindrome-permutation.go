// +build OMIT
// Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase
// that is sake forwards and backwards. A permutation is rearrangement of letters. The palindrome does not need to be
// limited to just dictionary words.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkPalindromePermutation("tact coa"))
	fmt.Println(checkPalindromePermutation("not a palindrome"))
}

func checkPalindromePermutation(s string) bool {
	charMap := make(map[rune]int)
	for _, c := range s {
		if c != ' ' {
			charMap[c] += 1
		}
	}

	maxOddChars := 0
	if len(strings.Replace(s, " ", "", -1)) % 2 != 0 {
		maxOddChars = 1
	}

	isPalindromePermutation := true
	oddCharsCount := 0
	for _, y := range charMap {
		if y % 2 != 0 {
			oddCharsCount++
		}

		if oddCharsCount > maxOddChars {
			isPalindromePermutation = false
			break
		}
	}

	return isPalindromePermutation
}