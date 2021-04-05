package main

import (
	"fmt"
)

func  main(){
	v:=5
	pv := &v
	fmt.Printf("v: %v\npv: %v\n",v,pv)

	*pv = 4
	fmt.Printf("v: %v\n",v)

	pv2 := &v
	fmt.Printf("pv: %v\npv2: %v",pv,pv2)
}