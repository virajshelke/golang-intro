package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// created a new reader to read from stdin i.e. console
	reader := bufio.NewReader(os.Stdin)

	// user prompt
	fmt.Print("Enter first number: ")
	// we call upon the ReadString method on reader to get the user input in string format
	// Note that the method returns 2 values & we are catching the data (first value) into input1
	// the second value is the error (if any). We are discarding/ignoring the error with '_'
	input1, _ := reader.ReadString('\n')

	// Second input
	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')

	// converting the string data to the float value (Also trimming the whitespaces before we convert)
	num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	num3 := num1 + num2
	fmt.Println("The addition of the two numbers is", num3)

}
