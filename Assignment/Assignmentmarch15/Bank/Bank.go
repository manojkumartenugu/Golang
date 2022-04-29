package main

import (
	// functions "Assignment/Assignmentmarch15/Bank/Functions"
	"fmt"

	functions "github.com/manojkumartenugu/Golangpractice/Assignment/Assignmentmarch15/Bank/Functions"
)

type bank interface {
	Account() int
}

func main() {

	x := functions.Deposit{TotalBalance: 150, DepositAmount: 10}

	y := functions.Withdraw{TotalBalance: 150, WithdawalAmount: 120}

	var b bank
	b = x
	fmt.Println("Total balance : ", b.Account())
	b = y
	fmt.Println("Total balance : ", b.Account())
}
