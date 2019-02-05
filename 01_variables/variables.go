package main

import (
	"fmt"
)

var globalVariable string = "My global variable"

func main() {
	// variable declaration
	var a string
	// variable initialization
	a = "variable initialization"

	// variable declaration and initialization
	var b string = "variable declaration and initialization"

	// sugar for declaration and initialization
	c := "sugar for declaration and initialization"

	// would produce an error, as this declares the variable again
	// c := "reassign c"

	// works as this doesn't declare c again
	c = "reassign c"

	// declare multiple variables at once
	var d, e string

	// declare and initialize multiple variables at once
	var (
		f bool   = false
		g string = "hey"
	)

	// define a constant
	const MY_CONSTANT int = 1
	// error, constants can't change their value
	// MY_CONSTANT = 2

	// print this shit  that the compiler doesn't complain
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(MY_CONSTANT)
	fmt.Println(globalVariable)
}
