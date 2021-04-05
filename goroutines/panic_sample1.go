package main

import "fmt"

func checkSafeNumber(numB int) int {
	if numB > 10 {
		panic(fmt.Sprintf("%v is pass the limit.", numB))
	} else {
		fmt.Printf("%v is a safe number.", numB)
	}
	return numB
}

func main() {
	var numBer int
	fmt.Print("Enter a number:")
	fmt.Scanln(&numBer)

	checkSafeNumber(numBer)
}
