package main 

import ( 
	"fmt"
) 

func main() { 
	var anon_func func(int) 
	anon_func = func(number int) { 
		if number == 0 { 
				return
		} else { 
			fmt.Println(number) 
			anon_func(number-1) 
		} 
	} 
	anon_func(5) 
} 
