package main

import (
	"fmt"
	"strconv"
)

const n = 20

func fizzbuzz(i int) string {
	switch {
	case i%15 == 0:
		return "fizzbuzz"
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	default:
		return strconv.Itoa(i)
	}
}

func main() {
	for i := 1; i < n; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
