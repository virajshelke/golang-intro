package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')

	num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	num3 := num1 + num2
	fmt.Println("The addition of the two numbers is", num3)

}
