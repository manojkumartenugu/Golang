package Withdraw

import "fmt"

func Withdraw(witcash uint, bal int) {

	fmt.Printf("your ammount is %v initially", bal)

	wit := bal - int(witcash)

	fmt.Println("your current ammount is:", wit)
	if wit < 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("your balance is", wit, "Please add the money")
			}
		}()
		panic("your do not have enough balance")
	}
}
