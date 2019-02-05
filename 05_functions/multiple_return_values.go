// Package main provides ...
package main

import "fmt"

// multiple return values,
// one of my favourite feature of golang
func getCoordinates() (int, int) {
	return 1, 2
}

func main() {
	// call the function and assign the values to x and y
	x, y := getCoordinates()
	// call  the function and throw away the second return value
	x, _ = getCoordinates()
	// call  the function and throw away the first return value
	_, y = getCoordinates()
	// call  the function and throw away both return values
	_, _ = getCoordinates()
	// the same as the one directly above
	getCoordinates()

	// would produce an error:
	// multiple-value getCoordinates() in single-value context
	// x = getCoordinates

	fmt.Println(x, y)
}
