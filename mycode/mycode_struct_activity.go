//Practice-101: Coding Hours
//Code Practice: Structs, Anonymous Struct
//code log: 1/10/2020
//programmer: veggiescubs
package main

import "fmt"

type FastFood struct {
	Price float64
	Meal  string
}

func main() {
	chicksilog := FastFood{
		78.90,
		"Chicken Sinangag with Egg",
	}
	fmt.Printf("Meal: %v", chicksilog)
}
