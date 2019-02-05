package main

import (
	"fmt"
)

func main() {
	if true {
		// would be executed
		fmt.Println("This is the if branch")
	}

	if false {
		// would not be executed
		fmt.Println("This is the if branch")
	} else {
		// would be executed
		fmt.Println("This is the else branch")
	}

	if false {
		// would not be executed
		fmt.Println("This is the if branch")
	} else if true {
		// would be executed
		fmt.Println("This is the else if branch")
	} else {
		// would not be executed
		fmt.Println("This is the else branch")
	}
}
