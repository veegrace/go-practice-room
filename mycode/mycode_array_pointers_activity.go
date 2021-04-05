//Practice-101: Coding Hours
//Code Practice: Array, Slice, Pointers
//code log: 1/10/2020
//programmer: veggiescubs

package main

import "fmt"

func main() {
	//declaring an initialized Array of Strings
	var arrayName = [8]string{"HONGJOONG", "SEONGHWA", "YUNHO", "YEOSANG", "SAN", "MINGI", "WOOYOUNG", "JONGHO"}
	fmt.Printf("arrayName: %v\n", arrayName) //print members of arrayName

	//getting a slice of an Array
	var sliceArray []string = arrayName[1:7]
	fmt.Printf("sliceArray: %v\n", sliceArray) //print slice of the arrayName

	//print all elements of arrays
	fmt.Println("ARRAY: ")
	for key, value := range arrayName {
		fmt.Printf("%v : %v\n", key, value)
	}

	//POINTERS
	pArrayName := &arrayName                   //pointer of arrayName is assigned to address of arrayName
	fmt.Printf("pArrayName: %v\n", pArrayName) //print members of pointer/address of arrayName

	//reassigning pointer pArrayName to a new Array value
	*pArrayName = [8]string{"BANGCHAN", "SEUNGMIN", "HAN", "CHANGBIN", "LEEKNOW", "FELIX", "HYUNJIN", "IN"}
	fmt.Printf("arrayName: %v\n", arrayName)
}
