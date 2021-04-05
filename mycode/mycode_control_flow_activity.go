package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Maria", "Seonghwa", "Yunho", "Lala","Mingi", "San", "Kung", "Poo" 
	}
	distribution = make(map[string] int, len(users))
)

func main(){
	coinsForUser := func(name string) int { //functions as variables param: string return: int
		var total int
		for i:=0; i<len(name); i++{ //loop for finding vowels for each name
			switch string(name[i]){ //control flow for vowels to each coin
			case "a", "A":
				total++
			case "e", "E":
				total++
			case "i", "I":
				total += 2
			case "o", "O":
				total += 3
			case "u", "U":
				total += 4
			}
		}
		return total
	}

	for _, name := range users {
		v := coinsForUser(name)
		if v > 10 {
			v=10
		}
		distribution[name] = v
		coins -= v
	}
	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}



