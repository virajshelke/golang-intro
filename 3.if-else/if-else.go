package main

import (
	"fmt"
)

func main() {
	x := 42

	// In go lang we don't use the parentheses
	if x < 0 {
		fmt.Println("Negative number")
	} else if x == 0 {
		fmt.Println("Number is 0")
	} else {
		fmt.Println("Positive number")
	}

	// we can declare temp variables for the contol structure
	// thus we can act temp on them and free up the space/mem (garbage collector will take care)
	// This allows us to further optimize the code
	if y := -10; y < 0 {
		fmt.Println("Negative number")
	} else if y == 0 {
		fmt.Println("Number is 0")
	} else {
		fmt.Println("Positive number")
	}

	// the variable y is not visible here i.e. it's out of scope here!
	// fmt.Println(y)

}
