package main

func add(x int, y int) int {
	return x + y
}

// add a loop in between to show different benchmark results
func add2(x int, y int) int {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}

	return x + y
}
