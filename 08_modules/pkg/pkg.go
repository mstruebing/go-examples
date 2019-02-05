package pkg

import "fmt"

// Capitalized is an exported function
func DoStuff() {
	fmt.Println("Doing stuff ...")
}

// same for constants/variables with capitalization
var MyConst string = "Some constant"

func Add(x int, y int) int {
	return x + y
}
