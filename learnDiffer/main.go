package main

import "fmt"

func main() {
	counting()
	defer fmt.Print("World\n")
	defer fmt.Print("Hello ")
	defer fmt.Println("One")
	fmt.Println("One")

}

func counting() {
	for i := range 5 {
		defer fmt.Printf("Value: %v\n", i)
	}
}
