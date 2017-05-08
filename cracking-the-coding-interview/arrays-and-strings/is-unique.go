// +build OMIT
// Implement and algorithm to determine if a string has all unique characters. What if you cannot use additional data
// structures?

package main

import "fmt"

const (
	// Considering ASCII chars only.
	maxNumberOfChars = 128
	uniqueString     = "abc"
	invalidString    = "test"
	emptyString      = ""
)

func main() {
	fmt.Println(isUnique(uniqueString))
	fmt.Println(isUnique(invalidString))
	fmt.Println(isUnique(emptyString))
}

func isUnique(str string) bool {
	switch {
	case len(str) > maxNumberOfChars:
		return false
	case len(str) < 2:
		return true
	default:
		checkboard := make([]bool, maxNumberOfChars)
		for _, char := range str {
			if checkboard[char] {
				// If it is true, then we see the same char for the second time.
				return false
			} else {
				checkboard[char] = true
			}
		}
		return true
	}
}
