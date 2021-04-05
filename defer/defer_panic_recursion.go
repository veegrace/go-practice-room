package main

import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker(){
	fmt.Println("about to panic")
	defer func(){
		if err := recover(); err != nil {
			log.Println("error: ", err)
		}
	}()
	panic("something bad happedened")
	fmt.Println("done panicking")
}