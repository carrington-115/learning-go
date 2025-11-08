package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// A simple rating app
	fmt.Println("Enter a value between 1 and 10 to rate this course: ")
	rateReader := bufio.NewReader(os.Stdin)
	rating, _ := rateReader.ReadString('\n')

	// add a system rate value of .5
	newRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	// error check
	if err != nil {
		fmt.Println("An error occurred durint conversions, ", err)
	} else {
		fmt.Println("Thank you for rating us ", newRating+0.5)
	}

}
