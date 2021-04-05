package main

import (
	"os"
	"fmt"
)

var ateezMembers = [] string {"HJ", "SH", "YH", "YS", "SN", "MG", "WY", "JH"}

func main(){
	file, err := os.Create("ateez.txt")

	if err != nil{ return }

	defer file.Close()
	for i:=0; i<len(ateezMembers);i++{
		file.WriteString(ateezMembers[i])
		file.WriteString("\n")
	}

	fmt.Println("file created.")

}