package main

import (
	"fmt"
)

func main() {
	c := '&'

	// this would fail with the default branch
	// delete the leading // to try this out
	// c = 'a'

	switch c {
	case '&':
		fmt.Println("&amp;")
	case '<':
		fmt.Println("&lt;")
	case '>':
		fmt.Println("&gt;")
	case '"':
		fmt.Println("&quot;")
	default:
		panic("unrecognized escape character")
	}
}
