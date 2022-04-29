package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Good Morning! Good Day!"
	substr := "Good"
	count := strings.Count(str, substr)
	fmt.Println("The number of occurrences of substring in the string is: ", count)
}
