package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----")

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		// doesn't execute the loop the fourth and fifth time
		if i == 3 {
			break
		}
	}

	fmt.Println("-----")

	for i := 0; i < 5; i++ {
		if i < 2 {
			// wouldn't print the value of i
			// when i < 2, so the first two loops
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("-----")

	// equivalent of a while loop
	sum := 1
	for sum < 10 {
		sum += sum
		// would be 16 at the end
	}

	var array = [5]int{10, 20, 30, 40, 50}

	// equivalent of a foreach loop
	for key, value := range array {
		fmt.Println(key, value)
	}

	// for {
	// 	fmt.Println("infinite ...")
	// }
}
