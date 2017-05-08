package main

import (
	"fmt"
)

func main() {
	fmt.Println(check("not", "anagram"))
	fmt.Println(check("test", "tset"))
	fmt.Println(check("joke", "ekoi"))
	fmt.Println(check("joke", "ekoj"))
}

func check(first, second string) bool {
	length := len(first)
	if length != len(second) {
		return false
	}

	for i := 0; i < length; i++ {
		if first[i] != second[length-i-1] {
			return false
		}
	}

	return true
}
