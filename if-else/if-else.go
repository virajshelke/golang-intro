package main

import (
	"fmt"
)

func main() {
	x := 42
	if x < 0 {
		fmt.Println("Negative number")
	} else if x == 0 {
		fmt.Println("Number is 0")
	} else {
		fmt.Println("Positive number")
	}

	if y := -10; y < 0 {
		fmt.Println("Negative number")
	} else if y == 0 {
		fmt.Println("Number is 0")
	} else {
		fmt.Println("Positive number")
	}

}
