package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers) // before

	// add 2 to every number

	// method: Using the old method
	for i := 0; i < len(numbers); i++ {
		numbers[i] += 2
	}
	fmt.Println(numbers) // after
	fmt.Println()

	// method 2: Using the range method
	for i := range numbers {
		fmt.Printf("%v: %v\n", i+1, numbers[i])
	}
	fmt.Println()

	// method 3: using the index and value
	for index, value := range numbers {
		fmt.Printf("%v: %v\n", index+1, value)
	}
	fmt.Println()

	// method 4: using the counting method
	var j uint16 = 0
	for j < 64 {
		if j == 10 {
			fmt.Println("The value of j is 5")
			break
		} else if j == 5 {
			goto someGoBetween
		}
		j++
	}

someGoBetween:
	fmt.Println("Skip the remaining code and come here")
}
