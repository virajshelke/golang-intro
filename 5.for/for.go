package main

import (
	"fmt"
)

func main() {

	// The following is an array of strings with the size of 3
	colors := [3]string{"Red", "Blue", "Green"}

	// basic for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println("Color is", colors[i])
	}

	// for loop with range
	// the first value which we discarded with '_' is the index value of the array element
	for _, color := range colors {
		fmt.Println("Color is", color)
	}

	// The following is an example of slice
	// Slices are better than array because we can dynamically add or remove elements & sort them efficiently
	colors_slice := []string{"Purple", "Violet"}

	// adding one more value to our slice using append function
	colors_slice = append(colors_slice, "Maroon")

	for _, color := range colors_slice {
		fmt.Println("Color is", color)
	}

	// The following is map declaration (key value pairs)
	countries := make(map[string]string)

	countries["India"] = "IN"
	countries["United States"] = "US"

	for key, value := range countries {
		fmt.Println(key, "country is", value)
	}

}
