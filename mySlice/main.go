package main

import "fmt"

func main() {
	fruits := []string{"mango", "banana", "orange"}
	var myAddedFruits = append(fruits, "tomato", "strawberry", "dragon fruit")
	var fruits1 = fruits[1:]
	var fruits2 = fruits[:2]
	fmt.Println(fruits, len(fruits))
	fmt.Println(myAddedFruits, len(myAddedFruits))
	fmt.Println(fruits1, len(fruits1))
	fmt.Println(fruits2, len(fruits2))

	// memory allocation in slices
	cities := make([]string, 3)
	cities[0] = "New York"
	cities[1] = "Yaounde"
	cities[2] = "London"
	// cities[3] = "London" // -> will give an index out of range error

	cities = append(cities, "Singapore", "Bangalore")
	fmt.Println("Cities with allocation: ", cities)

	// remove "London" from slice cities
	// "London" -> position [2] or 3

	var londonPosition int = 2
	cities = append(cities[:londonPosition], cities[londonPosition+1:]...)
	fmt.Println("Updated cities without London: ", cities)
}
