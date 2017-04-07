package main

import (
	"fmt"
	"math/rand"
)

const pi = 3.14

func swap(x, y int) (int, int) {
	return y, x
}

func display(x, y int) {
	fmt.Printf("%d and %d\n", x, y)
}

func duplicate(x, y int) (a, b int) {
	a = x * 2
	b = y * 2
	return
}

func main() {
	x := rand.Intn(10)
	y := -1
	display(x, y)

	x, y = swap(x, y)
	display(x, y)

	x, y = duplicate(x, y)
	display(x, y)

	var a, b float64 = 3.0, 5.0
	display(int(a), int(b))

	fmt.Printf("pi is of type %T\n", pi)

	fmt.Println(1 << 5)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	sum := 1
	for ; sum < 1000; { // This can be also "for sum < 1000 {".
		sum += sum
	}
	fmt.Println(sum)

	for {
		break
	}

	if n := 5; 5 > pi {
		fmt.Printf("%f is smaller than %d\n", pi, n)
	} else {
		fmt.Printf("%f is bigger than %d\n", pi, n)
	}

	switch test := "a"; test {
	case "a":
		fmt.Println("should be displayed")
		fallthrough
	case "b":
		fmt.Println("should be displayed")
	default:
		fmt.Println("shouldn't be displayed")
	}

	switch {
	case pi > 5:
		fmt.Printf("%f is bigger than 5\f", pi)
	}

	defer fmt.Println("the end")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
