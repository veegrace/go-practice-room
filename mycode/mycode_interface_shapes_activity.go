package main

import (
	"fmt"
	"math"
)

func main(){
	tr := 
}	

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Triangle struct{
	sideA, sideB, sideC float64
}

type Square struct{
	side float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}


//Heron's Formula : Sqrt(s(s-a)*(s-b)(s-c))
func (t Triangle) Area() float64{
	var s float64 = (t.sideA + t.sideB+ t.sideC) / 2
	return math.Sqrt(s*(s-t.sideA)*(s-t.sideB)*(s-t.sideC))
}

//Perimeter Formula: sideA + sideB + sideC
func (t Triangle) Perimeter() float64{
	return t.sideA + t.sideB + t.sideC
}

//Area Formula: s^2
func (sq Square) Area() float64{
	return sq.side * sq.side
}

//Perimeter Formula: 4s
func (sq Square) Perimeter() float64{
	return 4 * sq.side
}

//Area Formula: pi*r^2
func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

//Circumference: 2*Pi*r
func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

//Area Formula: l*w
func (r Rectangle) Area() float64{
	return r.length * r.width
}

//Perimeter Formula: 2l+2w
func (r Rectangle) Perimeter() float64{
	return 2 * (r.length + r.width)
}




