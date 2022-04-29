package main

import "fmt"

func panic(x int, y int) {
	fmt.Println(x / y)
}

func main() {
	panic(10, 0)

	var a, b int = 10, 0

	fmt.Println(a / b)

}
