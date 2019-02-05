package main

import "fmt"

// a simple add function with to arguments
// and an int as return value
func add(a int, b int) int {
	return a + b
}

func main() {
	// call the function
	result := add(1, 2)
	fmt.Println(result)
}
