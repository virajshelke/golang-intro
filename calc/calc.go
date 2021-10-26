package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	num1 := getInput()
	num2 := getInput()

	operation := getOperation()

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
	default:
		fmt.Println("Wrong operation selected!")
	}

}

func getOperation() byte {
	fmt.Print("Enter operation: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic("Error while reading the input from console")
	}
	operation := strings.TrimSpace(str)
	return operation[0]

}

func getInput() float64 {
	fmt.Print("Enter data: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic("Error while reading the input from console")
	}
	num, err2 := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err2 != nil {
		panic("Number was expected as an input")
	}
	return num
}
