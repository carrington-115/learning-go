package main

import "fmt"

func main() {
	// initial pointer
	var num *int
	fmt.Println("The number is ", num)
	fmt.Println("The number is ", &num)

	// definition
	var someNum = 20
	somePtr := &someNum
	fmt.Println("The number is ", *somePtr)
	fmt.Println("The address is ", somePtr)

	*somePtr += 2
	fmt.Println("New value is ", someNum)
}
