package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// user input and comma ok syntaxt
	/* Test with getting user data*/
	fmt.Println("Enter the user data: ")
	username := bufio.NewReader(os.Stdin)
	userEmail := bufio.NewReader(os.Stdin)
	userAge := bufio.NewReader(os.Stdin)
	userPassword := bufio.NewReader(os.Stdin)

	// extracting input with comma, ok syntax
	name, _ := username.ReadString('\n')
	email, _ := userEmail.ReadString('\n')
	age, _ := userAge.ReadString('\n')
	password, _ := userPassword.ReadString('\n')

	// printing user input
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Age: ", age)
	fmt.Println("Password: ", password)
}
