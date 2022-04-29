package Deposit

import (
	"fmt"
)

func DepositFun(cash uint, balance uint) {

	fmt.Println("your previouse balace is:", balance)

	var NewAmount uint

	if NewAmount <= 1000 {

		panic("the Maximum you can deposit is 1000")

	} else {

		NewAmount = cash + balance
	}
	fmt.Println("and your current balance is:", NewAmount)
}
