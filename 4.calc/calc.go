package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global reader variable
var reader = bufio.NewReader(os.Stdin)

// The following line gives error because we can't use the syntactic sugar here ':=' while declaring global var
// The ':=' is allowed inside functions
// reader := bufio.NewReader(os.Stdin)

func main() {
	num1 := getInput()
	num2 := getInput()

	operation := getOperation()

	// switch case in go lang don't need explicit break statements
	switch operation {
	case '+':
		fmt.Println("The addition of numbers is", num1+num2)
	case '-':
		fmt.Println("The subtraction of numbers is", num1-num2)
	case '*':
		fmt.Println("The multiplication of numbers is", num1*num2)
	case '/':
		fmt.Println("The division of numbers is", num1/num2)
		// fallthrough
		// uncomment the above line to let case '/' fallthrough and execute default as well
		// when case '/' is triggered.
		// This fallthrough behaviour is observed in C, C++ etc. & we use break to avoid it there
	default:
		fmt.Println("Wrong operation selected!")
	}

}

// function to get the value of the operation to be performed
// returns a character
func getOperation() byte {
	fmt.Print("Enter operation ('+', '-', '*', '/'): ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic("Error while reading the input from console")
	}
	operation := strings.TrimSpace(str)
	return operation[0]

}

// function to get number as input from user (from stdin)
func getInput() float64 {
	fmt.Print("Enter data: ")

	// Here we are interested in both the values i.e. data & error (if any)
	str, err := reader.ReadString('\n')

	// checking if we have any error
	if err != nil {
		// we can print the error or call upon the panic method which quits the program with proper exit code
		panic("Error while reading the input from console")
	}

	// converting string to float & checking error if any
	num, err2 := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err2 != nil {
		panic("Number was expected as an input! Couldn't convert string to number")
	}
	return num
}
