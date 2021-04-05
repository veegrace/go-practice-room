package main

import "fmt"

var (
	userYear, userMonth, userDay, userAge int
	currentYear                           int = 2020
)

func main() {
	fmt.Print("Enter year: ")
	fmt.Scanln(&userYear)

	userAge = currentYear - userYear
	fmt.Println("Your age: ", userAge)

	if userAge < 18 {
		fmt.Println("Child")
	} else {
		fmt.Println("Adult")
	}
}
