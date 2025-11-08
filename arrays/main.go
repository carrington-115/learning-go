package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println("The numbers in array<Numbers> are: ", numbers)
	fmt.Printf("\nThe type of array numbers is %T\n", numbers)

	// strings
	var names = [3]string{"Mark", "Carrington", "Fru"}
	fmt.Println("The names are: ", names)

	// lengths
	fmt.Println("Length of numbers: ", len(numbers))
	fmt.Println("Length of names: ", len(names))
}
