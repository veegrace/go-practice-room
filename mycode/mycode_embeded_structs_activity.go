//Practice-101: Coding Hours
//Code Practice: Structs, Embedding Structs
//code log: 1/10/2020
//programmer: veggiescubs

package main

import "fmt"

type Request struct{
	Resource string
}

type AuthenticatedRequest struct{
	Request
	Username, Password string
}

func main(){
	newRequest := new(AuthenticatedRequest){
		Request{"fb.com/request"},
		Username: "admin",
		Password: "P@ssw0rd",
	}

	fmt.Printf("Username: %v\nPassword: %v\nfrom %v\n", newRequest.Username,newRequest.Password,newRequest.Resource)
}

