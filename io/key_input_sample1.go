package main

import (
	"fmt"
)

var (
	name string
	
)

func main(){
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Printf("Your Name is %v.", name)
}
