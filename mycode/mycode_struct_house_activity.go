package main

import "fmt"

type House struct{
	NoRooms int
	Price float64
	City string
}

func main(){
	newHouse := House{
		NoRooms: 90,
		Price: 340.90,
		City: "Lapu-Lapu",
	}

	fmt.Println(newHouse.NoRooms)
}