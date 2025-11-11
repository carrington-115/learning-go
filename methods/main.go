package main

import "fmt"

func main() {
	NewUser := User{name: "Mark", email: "mark@mail.com"}
	name, email := NewUser.getUserDetails()
	fmt.Printf("My name is %v and my email is %v\n", name, email)
	newEmail := NewUser.setEmail("new-mark@mail.com")
	fmt.Println("My email has been change to ", newEmail)
}

type User struct {
	name  string
	email string
}

// getter
func (user User) getUserDetails() (string, string) {
	return user.name, user.email
}

// setter and getter
func (user User) setEmail(newEmail string) string {
	user.email = newEmail
	return user.email
}
