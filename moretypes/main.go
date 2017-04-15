package main

import (
	"fmt"
	"strings"
	"math"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printTicTacToe() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	var a [2]string // Array (constant size).
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4]) // Slice (dynamic size). It's just a reference to array!

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	slice := names[1:3]
	fmt.Println(slice)
	slice[0] = "XXX"
	fmt.Println(slice)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // Slice the slice to give it zero length.
	printSlice(s)

	s = s[:4] // Extend its length.
	printSlice(s)

	s = s[2:] // Drop its first two values.
	printSlice(s)

	s = make([]int, 5)
	printSlice(s)

	s = make([]int, 0, 5)
	printSlice(s)

	s = make([]int, 0)
	printSlice(s)

	s = append(s, 0, 5, 7, 8) // Append on nil slice.
	printSlice(s)

	printTicTacToe()

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	var m = map[string]Vertex{
		"Bell Labs": {
			34, 77,
		},
		"Google": {
			-12, 66,
		},
	}
	fmt.Println(m["Bell Labs"])

	delete(m, "Bell Labs")
	elem, ok := m["Bell Labs"]
	fmt.Println(elem, ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder() // Closures!
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
