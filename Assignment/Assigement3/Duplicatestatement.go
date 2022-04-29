package main

import (
	"fmt"
	"strings"
)

var statement1 string = "In this problem we are trying to find how many words are repeated in a sentence for example fear leads to angery angery leads to hatred hatred leads to confilct" //In this problem we are trying to find how many words are repeated in a sentence for example fear leads to angery angery leads to hatred hatred leads to confilct
var i, j int
var s []string
var word []string

func main() {

	fmt.Println(" the given sentence is", statement1)
	s = strings.Split(statement1, " ")
	for i = 0; i < len(s); i++ {
		for j = i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				word = append(word, s[i])
			}
		}

	}
	fmt.Println("the Duplicate words are as follows ", word)
	// fmt.Println(words)

}
