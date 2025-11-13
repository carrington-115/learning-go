package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Student  bool   `json:"is_student"`
}

func main() {
	fmt.Println("This is how JSON works")
	testCase := createJSON()
	fmt.Println(testCase)
}

func createJSON() string {
	users := []user{
		{"Mark", "mark@email.com", "mark2458", true},
		{"Carrington", "carrington@email.com", "omega285", false},
		{"Peter", "peter@email.com", "Pierre3892", true},
		{"Fru Mark", "fruchei@email.com", "mark2458", true},
	}
	// data, err := json.Marshal(users) // this gives no indents
	data, err := json.MarshalIndent(users, "", "\t")
	handleError(err)
	return string(data)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
