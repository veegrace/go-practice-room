package main

import "fmt"

type Numbers struct{
	X, Y int
	Symbol string 
}

func Sum(a, b int) int{
	return a + b
}

func Difference(a,b int) int{
	return a - b
}

func Product(a,b int) int {
	return a * b
}

func Calculator(n Numbers) int{
	Answer := 0

	switch {
	case n.Symbol == "+":
		Answer = Sum(n.X,n.Y)
	case n.Symbol == "-":
		Answer = Difference(n.X,n.Y)
	case n.Symbol == "*":
		Answer = Product(n.X,n.Y)
	default:
		fmt.Println("Invalid Equation.")
	}

	return Answer
}

func main(){
	num := Numbers{6,9,"*",}
	fmt.Println(Calculator(num))
}