package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	location := "D:/Golang/paths/test.txt"
	// paths := filepath.Join("dir", "subdir", "test.txt")
	// fmt.Println(paths)
	
	locate :=  filepath.Dir(location)
	base := filepath.Base(locate)
	dirs := filepath.Dir(locate)
	fmt.Println("base",base,"Dir",dirs)
	reads,_ := ioutil.ReadFile(location)
	// if err {
	// 	fmt.Errorf("eror")
	// }
	fmt.Println(string(reads))

	cleaned := filepath.Clean(location)
	fmt.Println("Clean: ", cleaned)
}
