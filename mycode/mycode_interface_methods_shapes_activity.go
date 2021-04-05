package main

import (
	"fmt"
	"math"
)

//main function
//test area
func main(){

	//usage of anonymous fields
	rec := Rectangle{5,4} //rectangle{legth,width}
	fmt.Println("Area of Rectangle: ", rec.Area())
	fmt.Println("Perimeter of Rectangle: ", rec.Perimeter())

	//data field of struct in Circle is declared
	cir := Circle{Radius: 2} //circle has radius value of 2

	//prints area and perimeter of a circle
	fmt.Println("Area of Circle: ", cir.Area())
	fmt.Println("Perimeter of Circle: ", cir.Perimeter())

	//interface types with concrete values
	var s Shape = Rectangle{5,4}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s,s)
	fmt.Println("Area of Rectangle: ", s.Area())
	fmt.Println("Perimeter of Rectangle: ", s.Perimeter())

	s = Circle{2}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s,s)
	fmt.Println("Area of Circle: ", s.Area())
	fmt.Println("Perimeter of Circle: ", s.Perimeter())

	totalArea := CalculateArea(Circle{2}, Rectangle{4,5}, Circle{3})
	fmt.Println("Total Area = ", totalArea)
}

//Shape interface 
type Shape interface{
	Area() float64 		//Area method
	Perimeter() float64 //Perimeter method
}


//rectangle struct
type Rectangle struct{
	Length, Width float64
}

//area of rectangle
//formula: a=l*w
func (r Rectangle) Area() float64{
	return r.Length * r.Width
}

//perimeter of rectangle
//formula: p=2l+2w
func (r Rectangle) Perimeter() float64{
	return 2*(r.Length+r.Width)
}


//Circle struct
type Circle struct{
	Radius float64
}

//Area of a Circle 
//formula: a=pi*r^2
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}


//Perimeter or Circumference of a circle
//formula: c=2*pi*r
func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}


//Variadic function to calculate the collection of area of shapes
func CalculateArea(shapes ...Shape) float64{	//returns total Area
	totalArea := 0.0	//initial value == 0.0
	for _, s := range shapes{	//for [key], [value]:= range [collectio]{...}
		totalArea =+ s.Area()	//totalArea = totalArea + s.Area()
	}
	return totalArea
}