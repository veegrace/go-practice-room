package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X,Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vertex{3,4} 	//new struct in an anonymous fields
					 	//setting X= 3, Y= 4
	fmt.Println(Abs(v))	//prints the sqrt of (3*3 + 4*4)=(9+16)=(25)
}