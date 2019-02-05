package main

import (
	"errors"
	"fmt"
)

// An error is a datatype
// often used as a second return value
func someFuncThatCanFail(shouldFail bool) (int, error) {
	if shouldFail {
		// create a new error with a message
		return 0, errors.New("something went wrong")
	}
	return 0, nil
}

func main() {
	// only use the error value
	var _, err = someFuncThatCanFail(false)

	// Yep, correct, thats how error handling in go works
	if err != nil {
		// would not be called
		// prints the error, in a real application you might want to
		// log this into a file or something
		fmt.Println(err)

		// do error handling
	}

	_, err = someFuncThatCanFail(true)
	if err != nil {
		// would be called
		fmt.Println(err)
		// do error handling
	}

	// a panic prints a message to stdout and exits the application
	// with an exit value != 0
	panic("THIS IS A PANIC AND EXITS THE APPLICATION IMMEDIATELY")
}
