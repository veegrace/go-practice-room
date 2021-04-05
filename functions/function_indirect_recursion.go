package main 

import ( 
	"fmt"
) 

func print_one(n int) { 
	
	if n >= 0 { 
		fmt.Println("In first function:", n) 
		print_two(n - 1) 
	} 
} 

func print_two(n int) { 

	if n >= 0 { 
		fmt.Println("In second function:", n) 
		print_one(n - 1) 
	} 
} 

func main() { 
	
	print_one(10) 
	print_one(-1) 
} 
