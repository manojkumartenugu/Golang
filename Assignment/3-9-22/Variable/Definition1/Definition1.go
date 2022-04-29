package main

import "fmt"

func main() {
	// varibale definition
	// varibale declaration
	// := is a shot hand form where we can define and declare varibale and value
	x := 10

	// Varibale definition
	var y int
	fmt.Println("Welcome to regions bank your current balance is", x, "How much would you like to deposit now:")

	fmt.Scan(&y)

	// varibale declaration
	x = x + y
	fmt.Println("Your current balance is:", x)

	// this is other type of defining and declaring which is not commonly use
	var z string = "Thank you"
	fmt.Println(z)

}
