// Description:
// You are given an array strarr of strings and an integer k.
// Your task is to return the first longest string consisting of k consecutive strings taken in the array.
//
// Example:
// longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
// n being the length of the string array, if n = 0 or k > n or k <= 0 return "".

package main

import "fmt"

func main() {
	fmt.Println(longestConsecutiveWords([]string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"}, 2))
}

func longestConsecutiveWords(strarr []string, k int) (result string) {
	if len(strarr) == 0 || k <= 0 || len(strarr) < k {
		return
	}

	maxLength := 0
	for i := 0; i < len(strarr) - k + 1; i++ {
		var word string
		for j:=0; j<k; j++ {
			word += strarr[i+j]
		}

		if len(word) > maxLength {
			maxLength = len(word)
			result = word
		}
	}

	return
}