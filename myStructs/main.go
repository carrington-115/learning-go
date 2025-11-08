package main

import "fmt"

func main() {
	newUser := User{"mark", "324308ldfksl", "21-03-2006", 15}
	fmt.Println("New user built with struct: ", newUser)
	fmt.Printf("New user built with struct: %+v\n", newUser)

	// presentation
	fmt.Printf("My name is %v and I am %v years old.\n", newUser.name, newUser.age)

}

type User struct {
	name string
	enId string
	dob  string // DD-MM-YYYY
	age  int
}
