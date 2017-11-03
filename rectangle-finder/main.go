// The task is to detect whether a given 2D array (containing 1's and O's) contains a rectangle.
// Rectangle is being defined as having 1's at all its 4 corners.
//
// For example, given a following array with N rows and M cols:
//
// O 1 O 1 O O
// O O 1 O 1 O
// O O O O 1 O
// O O 1 O 1 O
//
// The function should return true.

package main

import (
	"fmt"
)

func containsRectangle(array [][]bool) bool {
	// Creates map of 1's positions in each row.
	var cornersPositions [][]int = make([][]int, 0)
	for _, row := range array {
		cornersPositions = append(cornersPositions, getCornerPositions(row))
	}

	var i, j int
	for i = 0; i < len(cornersPositions); i++ {
		for j = i + 1; j < len(cornersPositions); j++ {
			if countDuplicates(cornersPositions[i], cornersPositions[j]) >= 2 {
				return true
			}
		}
	}

	return false
}

func getCornerPositions(array []bool) (positions []int) {
	for index, element := range array {
		if element == true {
			positions = append(positions, index)
		}
	}
	return
}

func countDuplicates(a, b []int) (duplicateCount int) {
	mapA := make(map[int]bool)
	for _, element := range a {
		mapA[element] = false
	}

	for _, element := range b {
		if _, ok := mapA[element]; ok {
			duplicateCount++
		}
	}
	return
}

func main() {
	example := [][]bool{
		{false, true, false, true, false, false},
		{false, false, true, false, true, false},
		{false, false, false, false, true, false},
		{false, false, true, false, true, false},
	}

	fmt.Println(containsRectangle(example))
}
