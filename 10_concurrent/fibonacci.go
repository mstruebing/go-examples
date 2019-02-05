package main

import (
	"fmt"
)

func fibonacci(n int) {
	a := 0
	b := 1
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
	}

	fmt.Println(a)
}

func main() {
	fibonacci(10000000000)

	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)
	// go fibonacci(10000000000)

	// fmt.Scanln()
}
