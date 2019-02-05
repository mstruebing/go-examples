package main

import (
	"fmt"
)

// We need to declare a type Convert
// which is a function which takes a string as an argument
// and returns a string
type Convert func(string) string

// first convert function
func ask(smth string) string {
	return fmt.Sprintf("%s????", smth)
}

// second convert function
func exclamate(smth string) string {
	return fmt.Sprintf("%s!!!!", smth)
}

// convert takes a string and a function which is of type Convert
func convert(smth string, fn Convert) string {
	return fmt.Sprintf("%s", fn(smth))
}

func main() {
	fmt.Println(convert("what", ask))       // produces "what???"
	fmt.Println(convert("what", exclamate)) // produces "what!!!"
}
