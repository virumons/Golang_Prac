package main

import (
	"fmt"
)

type engines struct{
	rpm uint16
	mpn uint16
	ownerinfo owener
}
type owener struct{
	name string
}

func (e engines) milestone() uint16{
	return e.rpm*e.mpn
}
func milesleft(e engines, miles uint16){
	if miles < e.milestone(){
		fmt.Println("can be reached")
		fmt.Println("",miles,e.milestone())
	}else{
		fmt.Println("cannot be reached")
	}
}
func main(){
	fmt.Println("hello")
	var e1 engines = engines{23,56,owener{ "John Doe"}}
	fmt.Println(e1)
	milesleft(e1,50)
}