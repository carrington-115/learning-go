package main

import "fmt"

func main() {
	// map for cities and their codes
	cities := make(map[string]string)
	cities["NY"] = "New York"
	cities["YD"] = "Yaounde"
	cities["BLR"] = "Bengaluru"
	cities["HK"] = "Hong Kong"

	fmt.Println("Map of cities: ", cities)
	delete(cities, "YD")
	fmt.Println("Map of cities: ", cities)

	// looping over maps
	for key, value := range cities {
		fmt.Printf("%v: %v\n", key, value) // key: value
	}
}
