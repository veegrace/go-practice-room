package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	//instance of Person
	hj := new(Person)

	//original
	hj.Name = "Kim Hongjoong"
	hj.Age 	=  22
	fmt.Printf("Name: %v\nAge: %v\n",hj.Name,hj.Age)

	//copies the instance of hj
	sh := hj //sets copy
	sh.Name = "Park Seonghwa" //renaming copied values
	sh.Age 	= 22
	fmt.Printf("Name: %v\nAge: %v\n",hj.Name,hj.Age)

	yh := *sh
	yh.Name = "Jeong Yunho"
	yh.Age	= 21
	fmt.Printf("Name: %v\nAge: %v\n",hj.Name,hj.Age)
	
	fmt.Printf("Name: %v\nAge: %v\n",hj.Name,hj.Age)
	fmt.Printf("Name: %v\nAge: %v\n",sh.Name,sh.Age)
	fmt.Printf("Name: %v\nAge: %v\n",yh.Name,yh.Age)

}