package main

import (
	"fmt"
	"regexp"
)

func main(){
	pattern := regexp.MustCompile(`^a[a-z]+b$`)

	teststr := []string{"acb","akjbdb","sjhas"}

	for _,str := range teststr{
		if pattern.MatchString(str){
			fmt.Printf("%s Matches the pattern\n",str)
		}else{
			fmt.Printf("%s does not matches to pattern \n",str)
		}
	}
}