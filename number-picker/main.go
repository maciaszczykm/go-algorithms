// Write a function that is able to pick all the numbers that are not in given collection and their sum is not bigger
// than given number.

package main

import (
	"fmt"
)

func quicksort(array []int32) []int32 {
	if len(array) < 2 {
		return array
	}

	left := 0
	right := len(array) - 1
	pivot := left + ((right - left) / 2)
	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	quicksort(array[:left])
	quicksort(array[left + 1:])

	return array
}

func pickNumbers(collection []int32, sum int32) (result []int32) {
	if len(collection) < 1 || sum < 1 {
		return
	}

	// Map usage would allow to skip sort, so it is a better solution.
	quicksort(collection)

	var number int32 = 1
	var collectionIdx int = 0
	var collectionNumber int32 = collection[collectionIdx]
	var isCollectionChecked = false

	for ; number <= sum; number++ {
		if collectionNumber == number && !isCollectionChecked{
			collectionIdx++
			if collectionIdx >= len(collection) {
				isCollectionChecked = true
			} else {
				collectionNumber = collection[collectionIdx]
			}
		} else {
			result = append(result, number)
			sum -= number
		}
	}

	return
}

func main() {
	fmt.Println(pickNumbers([]int32{1, 3, 4}, 7))
}
