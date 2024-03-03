package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello viraj")
	prime()
	loops()
}

// loops function 
func prime(){
	i:=1
	for i <= 10{
		if i%2 != 0{
			fmt.Println(i)
		}
		i++
	}
}

func loops(){
	var i int
	for i=0; i<=10 ; i++{
		fmt.Println(i)
	}
}