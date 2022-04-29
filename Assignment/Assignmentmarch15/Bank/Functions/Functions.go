package functions

import "fmt"

type Deposit struct {
	TotalBalance  int
	DepositAmount int
}
type Withdraw struct {
	TotalBalance    int
	WithdawalAmount int
}

func (d Deposit) Account() int {
	d.TotalBalance = d.TotalBalance + d.DepositAmount
	return d.TotalBalance
}

func (w Withdraw) Account() int {
	if w.TotalBalance > 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("your balace is less please add money")
			}
		}()
		panic("you do not have enough balance")
		//w.TotalBalance = w.TotalBalance - w.WithdawalAmount
	}
	return w.TotalBalance
}
