package main

import "fmt"

const LoginUrl string = "/api/v1/auth/login"

func main() {
	// int, string, boolean
	var age int = 22
	var name string = "Mark"
	fmt.Println("My name is ", name, " and I am ", age, " years old")
	fmt.Printf("Mark is of type %T \n", name)
	var likesAppleSauce bool = false
	var process_speed int = 400
	var time_delay uint16 = 140
	// var time_delay2 uint8 = 500
	fmt.Println("Process speed and time delay ", process_speed, " ", time_delay, " ", likesAppleSauce)

	// explicit variables
	var newName = "Carrington"
	var score = 20
	var population = 230083
	var isHandsome = true
	driverStatus := "available"
	fmt.Printf("\nTypes: %T, %T, %T, %T", newName, score, population, isHandsome)
	fmt.Println(LoginUrl, driverStatus)
}
