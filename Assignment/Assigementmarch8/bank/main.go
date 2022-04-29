package main

import (
	// "Assignment/Assigementmarch8/bank/Deposit"
	// "Assignment/Assigementmarch8/bank/Withdraw"
	"fmt"
	"github.com/manojkumartenugu/Golangpractice/Assignment/Assigementmarch8/bank/Deposit"
	"github.com/manojkumartenugu/Golangpractice/Assignment/Assigementmarch8/bank/Withdraw"
)

func main() {
	fmt.Println("welcome to the bank system")
	fmt.Println("your ammount is 100 initially")

	fmt.Println("Deposit or Withdraw")
	var dw string
	fmt.Scanln(&dw)
	var amou uint
	fmt.Println("Enter the ammount")
	fmt.Scanln(&amou)

	if dw == "Deposit" || dw == "deposit" {
		Deposit.DepositFun(amou, 1000)
		//Deposit.DepositFun(amou, 1000)
	} else if dw == "Withdraw" || dw == "withdraw" {
		//Withdraw.Withdraw(amou, 1000)
		Withdraw.Withdraw(amou, 1000)
	}

}
