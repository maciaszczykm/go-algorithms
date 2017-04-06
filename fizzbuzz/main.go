package main

import "fmt"

const n = 20

func main() {
	for i := 1; i < n; i++ {
		if i % 15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}