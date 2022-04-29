package main

import "fmt"

// here in this program we will see what is a variable definition and declaration the diffrence and how to use them (:=,=, definition, declaration)

// func futureage(i int, j int){

// }

func main() {

	fmt.Print("What is your current age:")
	var x int
	fmt.Scanln(&x)
	fmt.Print("your current age is:", x, " your want your age after how many years:")
	var y int
	fmt.Scanln(&y)
	// fmt.Println(x + y)

	// var z int
	z := x + y

	fmt.Printf("Your age after %v years will be %v", y, z)
	fmt.Println("")

}
