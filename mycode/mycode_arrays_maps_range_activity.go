package main

import "fmt"

var (
	numberArray = []int{13,14,15,29,24,10}
	stringArray = []string{"HJ", "SH", "YH", "YS", "SN", "MG", "WY", "JH"}
	sum int = 0
)

func main(){
	for _, value := range numberArray{
		sum += value
	} 
	fmt.Println(sum)

	for key, value := range stringArray{
		fmt.Printf("%v : %v\n", key, value)
	}
}