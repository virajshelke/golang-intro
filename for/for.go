package main

import (
	"fmt"
)

func main() {

	colors := [3]string{"Red", "Blue", "Green"}

	for i := 0; i < len(colors); i++ {
		fmt.Println("Color is", colors[i])
	}

	for _, color := range colors {
		fmt.Println("Color is", color)
	}

	countries := make(map[string]string)

	countries["India"] = "IN"
	countries["United States"] = "US"

	for key, value := range countries {
		fmt.Println(key, "country is", value)
	}

}
