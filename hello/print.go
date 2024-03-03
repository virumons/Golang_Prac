package main
// cmd
// go mod init filename 
// the command create a mod file, for dependecies setup and create module path for the file

// check output
// go run . 
// or
// go run filename.go

// variables 
// int -> has -> int16,int32,int64
// CANT HAVE ARITHMETIC OPERATIONS ON DIFF DATATYPES like - float+int  

// format specifiers
// %v - default format for the value.
// %T - prints the type of the value.
// %d - decimal integer.
// %b - binary representation.
// %o - octal representation.
// %x, %X - hexadecimal representation (lowercase/uppercase).
// %f - floating-point numbers.
// %e, %E - scientific notation (lowercase/uppercase).
// %s - string.
// %p - pointer, printed in hexadecimal with a leading 0x.
// %t - boolean, the word true or false.

import (
    "errors"
    "fmt"
)

func main() {
    var number int = 30
    arr := [3]int{1,2,3}
    fmt.Println(arr)
    fmt.Println(number)
    printme("hello world",50)
    fmt.Println("Hello, World!")

}

func printme(printno string, number int) error {
    var err error
    if number == 50{
        fmt.Printf("got 50")
    }
    err= errors.New("cannot do")
    fmt.Println(printno + "\n");
    return err   
}

// err.Error() gives you an actual description of what went
// err is a variable  that holds an error value.
// to return multiple values from same function
// func multireturn(input value) (int,string){
    // return R1,R2
// }

// array decalration 
// var arr [5]int   // Declaring and initializing an array

 
