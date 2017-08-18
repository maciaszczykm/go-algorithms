// +build OMIT
// Write a method to replace all spaces in a string with `%20`. You may assume that the string has sufficient space
// at the end to hold the additional characters, and that you are given the "true" length of the string.

package main

import "fmt"

var escapedSpace = []byte{'%', '2', '0'}

func main() {
	fmt.Println(urlify("Mr John Smith    ", 13))
}

func urlify(s string, l int) string {
	spaceDiff := len(s) - l
	url := []byte(s)
	for i := l - 1; i >= 0; i-- {
		if url[i] != byte(' ') {
			url[i+spaceDiff] = url[i]
		} else {
			url[i+spaceDiff-2] = escapedSpace[0]
			url[i+spaceDiff-1] = escapedSpace[1]
			url[i+spaceDiff] = escapedSpace[2]
			spaceDiff -= 2
		}
	}

	return string(url)
}
