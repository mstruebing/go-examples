package main

import "fmt"

// variadic functions can take
// as many arguments of the same type as you like
func myVariadicFunc(ints ...int) {
	fmt.Println(ints)
}

func sum(ints ...int) int {
	result := 0

	// loop over the int array
	// and throw away the key (which would be the array index)
	for _, value := range ints {
		result += value
	}

	return result
}

func main() {
	// pass in as many arguments as you want of the same type for
	// variadic functions
	myVariadicFunc(1, 2, 3)
	myVariadicFunc(1, 2, 3, 4, 5)

	fmt.Println(sum(1, 2, 3, 4, 5))
}
