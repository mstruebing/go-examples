package main

import "fmt"

func main() {
	// inline function aka anonymous function
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	func() {
		fmt.Println("HELLO WORLD")
	}()
}
