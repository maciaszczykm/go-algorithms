package main

import "fmt"

var (
	original = map[string]string{"1": "abc", "2": "xyz", "3": ""}
	changed  = map[string]string{"1": "abc", "3": "mod", "4": "new"}
)

func main() {
	compare(original, changed)
}

func compare(first, second map[string]string) {
	for key, value := range first {
		secondValue, ok := second[key]
		switch {
		case !ok:
			fmt.Printf("DEL  %v : %v\n", key, value)
		case ok && value == secondValue:
			fmt.Printf("OLD %v : %v\n", key, value)
		case ok && value != secondValue:
			fmt.Printf("MOD  %v : %v\n", key, value)
		}
	}

	for key, value := range second {
		_, ok := first[key]
		switch {
		case !ok:
			fmt.Printf("ADD  %v : %v\n", key, value)
		}
	}
}
