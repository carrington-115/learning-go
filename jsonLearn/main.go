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

	// decoding
	fmt.Println("\nDECODING THE JSON WEB DATA")
	var fromWeb string = `
		{
    	    "name": "Peter",
    	    "email": "peter@email.com",
    	    "is_student": true
    	}
	`
	decodeJSON(fromWeb)
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

func decodeJSON(data string) {
	var uniqueUser user
	webData := []byte(data)
	var isValid bool = json.Valid(webData)

	if isValid {
		fmt.Println("The JSON data from the web is valid type")
		json.Unmarshal(webData, &uniqueUser)
		fmt.Println(uniqueUser)
	} else {
		fmt.Println("The Data was not of valid JSON type")
	}

	// populating in another way
	var unpopulatedData map[string]interface{}
	json.Unmarshal(webData, &unpopulatedData)
	fmt.Println("Decoded unpopulated data: ")
	fmt.Println(unpopulatedData)

	// looping
	for key, val := range unpopulatedData {
		fmt.Printf("Key is %v, Value is %v, Value type is %T\n", key, val, val)
	}

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
