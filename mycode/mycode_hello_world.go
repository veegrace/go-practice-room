//source: golangr.com
//Begginer exercise
//Hello World
//Create a program that shows your name

package main

import "fmt"

var (
	myName string = "Verlin"
	myAddress string = "Suba Masulog"
)


func main(){
	fmt.Printf("Hello, %v. Your current address is %v", myName, myAddress)
}