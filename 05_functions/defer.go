package main

import "fmt"

// defer is always executed at the end of the function
// so you can not forget to close some buffers, files or connections
func myDeferFunc() int {
	defer fmt.Println("This should be printed at the end")
	// do very complex stuff again
	fmt.Println("complex stuff completed")
	return 1
}

func main() {
	fmt.Println(myDeferFunc())
}
