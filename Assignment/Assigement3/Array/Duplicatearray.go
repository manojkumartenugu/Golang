package main

import (
	"fmt"
)

var arrayvalue = []int{2, 3, 4, 5, 6, 6, 7, 7, 8, 8, 2, 4}
var i, j int
var newarrayvalue []int

func main() {
	fmt.Println("Find Duplicate number in a array")

	fmt.Printf("The given list of number as %v\n", arrayvalue)

	for i = 0; i < len(arrayvalue); i++ {
		for j = i + 1; j < len(arrayvalue); j++ {
			if arrayvalue[i] == arrayvalue[j] {
				newarrayvalue = append(newarrayvalue, arrayvalue[i])
			}

		}
	}
	fmt.Println(newarrayvalue)
}
