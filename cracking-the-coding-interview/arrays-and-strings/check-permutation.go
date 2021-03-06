// +build OMIT
// Given two strings, write a method to decide if one is a permutation of the other.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(checkPermutation("marcin", "rnicam"))
	fmt.Println(checkPermutation("abc", "xyz"))
	fmt.Println(checkPermutation("abc", "abcd"))
	fmt.Println(checkPermutation("", "asd"))
}

func checkPermutation(a, b string) bool {
	// Sort is easy, but creating array of letters for each word and comparing them is also a good option.
	return len(a) == len(b) && sortString(a) == sortString(b)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
