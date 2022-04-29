package main

import (
	"fmt"
	"sort"
)

/*Lonely numbers problem:
Given an array of numbers, find the lonely numbers
 (A number which does not have matching number(eg: 3 & 3), or no consequtive numbers(eg: 1,2,3 are consequtive numbers)).*/

func main() {

	a := []int{2,1, 8, 9, 5, 4, 12}

	fmt.Println(lonelyNumber(a))

}

func lonelyNumber(a []int) []int {
	if len(a) < 2 {
		return a
	}

	sort.Ints(a)
	l := make([]int, 0, 1024)

	if a[0]+1 < a[1] {
		l = append(l, a[0])
	}
	if a[len(a)-1] > a[len(a)-2]+1 {
		l = append(l, a[len(a)-1])
	}
	for i := 1; i < len(a)-1; i++ {
		if a[i-1]+1 < a[i] && a[i]+1 < a[i+1] {
			l = append(l, a[i])
		}
	}
	return l
}
