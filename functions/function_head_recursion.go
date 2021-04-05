package main 

import ( 
	"fmt"
) 
func print_num(n int) { 
	if n > 0 { 
		print_num(n-1) 
		fmt.Println(n) 
	} 
} 

func main() { 
	print_num(5) 
} 
