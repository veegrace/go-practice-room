package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		fmt.Println("Good Morining")
	case currentTime.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evevning")
	}
	fmt.Println(currentTime)
	fmt.Printf("%v : %v", currentTime.Hour(), currentTime.Minute())
}
