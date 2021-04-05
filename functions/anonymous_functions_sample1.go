package main

import "fmt"

func main(){
	var divide func(float64, float64) (float64, error)
	divide = func(dividend, divisor float64) (float64, error){
		if divisor == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return dividend / divisor, nil
		}
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}