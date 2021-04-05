package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main(){
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1,7)
	fmt.Printf("Random Number: %v", randomNum)
}