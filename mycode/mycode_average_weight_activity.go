package main

import "fmt"

func main() {
	initial, weight, average, people := 0, 0, 0, 0
	fmt.Print("Enter number of people: ")
	fmt.Scanln(&people)
	for i := 1; i < people; i++ {
		fmt.Print("Enter weight ")
		fmt.Scanln(&weight)
		initial += weight
	}
	average = initial / 2
	fmt.Println("Total Weight: ", initial)
	fmt.Println("Average Weight: ", average)
}
