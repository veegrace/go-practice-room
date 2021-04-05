package main

import (
	"fmt"
)

type Ateez struct { //field names - type
	number     int
	leaderName string
	companions []string
}

func main() {
	aAteez := Ateez{
		number:     8,
		leaderName: "Kim Hongjoong",
		companions: []string{
			"Park Seonghwa",
			"Yunho Jeong",
			"Yeosang Kang",
			"Choi San",
			"Song Mingi",
			"Jung Wooyoung",
			"Choi Jongho",
		},
	}
	fmt.Println(aAteez)
	fmt.Println(aAteez.leaderName)
	fmt.Println(aAteez.number)
	fmt.Println(aAteez.companions)
	fmt.Println(aAteez.companions[0])
}
