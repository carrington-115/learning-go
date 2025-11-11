package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculator: ")
	fmt.Print("Enter calculator operation:\t")

	// calculator operation
	operationReader := bufio.NewReader(os.Stdin)
	operation, _ := operationReader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	// operation parameters
	fmt.Print("\nEnter 2 numbers that you want to add:\n")
	fmt.Print("Number 1: ")
	readerOne := bufio.NewReader(os.Stdin)
	n1s, _ := readerOne.ReadString('\n')
	n1s = strings.TrimSpace(n1s)
	n1, err1 := strconv.Atoi(n1s)

	fmt.Println("Number 2: ")
	readerTwo := bufio.NewReader(os.Stdin)
	n2s, _ := readerTwo.ReadString('\n')
	n2s = strings.TrimSpace(n2s)
	n2, err2 := strconv.Atoi(n2s)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number input")
		return
	}

	sysCalculator := Calculator{
		add:      add,
		subtract: subtract,
		divide:   divide,
		multiply: multiply,
	}

	// operation cases
	switch operation {
	case "+":
		fmt.Printf("%v + %v = %v\n", n1, n2, sysCalculator.add(n1, n2))
	case "-":
		fmt.Printf("%v - %v = %v\n", n1, n2, sysCalculator.subtract(n1, n2))
	case "*":
		fmt.Printf("%v * %v = %v\n", n1, n2, sysCalculator.multiply(n1, n2))
	case "/":
		divideResult := sysCalculator.divide(n1, n2)
		if divideResult == -1 {
			fmt.Println("Division by zero is not possible")
		} else {
			fmt.Printf("%v / %v = %v\n", n1, n2, sysCalculator.divide(n1, n2))
		}
	default:
		fmt.Println("Something went wrong. Try again!")
	}
}

// calculator type
type Calculator struct {
	add      func(n1 int, n2 int) int
	subtract func(n1 int, n2 int) int
	divide   func(n1 int, n2 int) int
	multiply func(n1 int, n2 int) int
}

// functions
func add(n1 int, n2 int) int {
	return n1 + n2
}

func subtract(n1 int, n2 int) int {
	return n1 - n2
}

func multiply(n1 int, n2 int) int {
	return n1 * n2
}

func divide(n1 int, n2 int) int {
	if n2 != 0 {
		return n1 / n2
	}
	return -1
}
