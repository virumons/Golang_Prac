package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file handling")

	// reading the files
	content, _ := ioutil.ReadFile("text.txt")

	fmt.Println(string(content))

	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Printf("error")
	}
	_, err = file.WriteString("Go lang")
	if err != nil {
		fmt.Printf("error in writing to a file %v", err)
	}
	fmt.Println("Data string written successfully")
	defer file.Close() //closing the connection after use

	res, _ := ioutil.ReadFile("text.txt")
	fmt.Println(string(res))

	modify := []byte(string(content)+"hello from me")

	ioutil.WriteFile("text.txt",modify,0644) 
	
}
