package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time: ")
	myTime := time.Now()
	fmt.Println(myTime)

	// time formatting
	fmt.Println("Date format: ", myTime.Format("01-02-2006"))
	fmt.Println("Time format: ", myTime.Format("15:04:05"))
	fmt.Println("Day format: ", myTime.Format("Monday"))
	fmt.Println("Month format: ", myTime.Format("Jan"))
	fmt.Println("Day format: ", myTime.Format("2"))

	// Create my own date
	myDate := time.Date(2026, time.March, 18, 6, 0, 0, 0, time.UTC)
	fmt.Println("My date: ", myDate)
	fmt.Println("Custom date: ", myDate.Format("01-02-2006"))
	fmt.Println("Custom month: ", myDate.Format("Jan"))
	fmt.Println("Custom day: ", myDate.Format("Monday"))
	fmt.Println("Custom time: ", myDate.Format("15:04:05"))
}
