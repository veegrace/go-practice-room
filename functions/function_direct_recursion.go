package main 

import ( 
	"fmt"
) 

func factorial_calc(number int) int { 
 
	if number == 0 || number == 1 { 
		return 1 
	} 
 
	if number < 0 { 
		fmt.Println("Invalid number") 
		return -1 
	} 

	return number*factorial_calc(number - 1) 
} 

func main() { 

	answer1 := factorial_calc(0) 
	fmt.Println(answer1, "\n") 
	
	answer2 := factorial_calc(5) 
	fmt.Println(answer2, "\n") 
	
	answer3 := factorial_calc(-1) 
	fmt.Println(answer3, "\n") 
	
	answer4 := factorial_calc(6) 
	fmt.Println(answer4, "\n") 
} 
