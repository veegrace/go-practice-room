//Practice-101: Coding Hours
//Code Practice: Pointers(Deferencing), Structs
//code log: 1/10/2020
//programmer: veggiescubs

package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//constructing a new struct type using dot notation
	naruto := new(Person)
	naruto.Name = "Naruto Uzamaki"
	naruto.Age = 33
	naruto.Address = "Konohagakure"
	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", naruto.Name, naruto.Age, naruto.Address)

	boruto := naruto
	boruto.Name = "Boruto Uzamaki"
	boruto.Age = 14
	boruto.Address = "Hidden Leaf Village"
	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", boruto.Name, boruto.Age, boruto.Address)
	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", naruto.Name, naruto.Age, naruto.Address)

	//deferencing pointers
	hinata := *boruto
	hinata.Name = "Hinata Uzamaki"
	hinata.Age = 30
	hinata.Address = "CoHo"

	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", naruto.Name, naruto.Age, naruto.Address)
	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", boruto.Name, boruto.Age, boruto.Address)
	fmt.Printf("Name: %v\nAge: %v\nAddress: %v\n", hinata.Name, hinata.Age, hinata.Address)
}
