package main

import (
	"fmt"
)

func main() {
	var things1 = []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", things1)
	result := square(&things1)
	fmt.Println("the memory location", &things1)
	fmt.Println("Squared slice:", result)
}

func square(things2 *[]int) []int {
	for i := range *things2 {
		(*things2)[i] = (*things2)[i] * (*things2)[i]
	}
	return *things2
}
