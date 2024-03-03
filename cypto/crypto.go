package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("Maths,Crypto,Random")

	// var num int = 2
	// var num2 float64 = 4

	// fmt.Println("The sum :", num+int(num2))

	// Random  Number Generation
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//  Generate a random number via crypto
	// myran, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(myran)
	
	gensha("hello world")
}

func gensha(s string) {
	input := s
	// hash calculate
	hash := sha256.Sum256([]byte(input))
	// convert to hexadecimal format
	result := fmt.Sprintf("%x", hash)
	fmt.Printf("\n The SHA-256 of %q is %s: \n\n", input, result)

	hashstr := result

	constr,err := hex.DecodeString(hashstr)
	if  err != nil{
		fmt.Errorf("error decoding the string %v", err)
	}

	// hashedstr := string(constr)
	fmt.Println("Decoded hash to String",string(constr))
}