package main

import (
	"fmt"
	"sync"
	//"time"
)

var wg = sync.WaitGroup{}

func main(){
	var msg = "Hello"
	wg.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
}
