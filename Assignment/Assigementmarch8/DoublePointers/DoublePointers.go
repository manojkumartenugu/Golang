package main

import (
	"fmt"
)

func main() {
	x := 100
	var y *int = &x  //using &x giving value of x location to y and refering the value by using *
	var z **int = &y // using &y giving value to z of location y and refering to

	// & is used to get the location where the x value is stored
	// * is use to for refernce

	fmt.Println("The value in the variable x is", x)
	fmt.Println("Memory address of variable x is (pulling the value by y value)", y)
	fmt.Println("Memory address of variable x is (pulling the value by using pointer(&)to x)", &x)
	fmt.Println()
	fmt.Println("The value in the variable y is", y)
	fmt.Println("The refernce for value y is ", *y)
	fmt.Println("Memory address of variable y is (pulling the value by z value)", z)
	fmt.Println("Memory address of variable y is (pulling the value by using pointer(&)to y)", &y)
	fmt.Println()
	fmt.Println("The value in the variable z is", z)
	fmt.Println("The refernce for value z is ", *z)
	fmt.Println("The double Pointer z is ", **z) // using double refernce here
	fmt.Println("Memory address of variable y is (pulling the value by using pointer(&)to z)", &z)

}
