package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-oop/accounts"
	"github.com/LeonardoCampello-dev/go-oop/customers"
)

type verifyAccount interface {
	Withdraw(value float64) (string, float64)
}

func payBankSlip(account verifyAccount, bankSlipValue float64) {
	account.Withdraw(bankSlipValue)
}

func main() {
	leonardo := customers.Holder{
		Name:           "Leonardo",
		DocumentNumber: "780.219.590-02",
		Profession:     "Developer",
	}

	savingsAccount := accounts.SavingsAccount{
		Holder:  leonardo,
		Agency:  111,
		Account: 1,
	}

	savingsAccount.Deposit(5000)

	payBankSlip(&savingsAccount, 900)

	fmt.Println(savingsAccount.GetBalance())
}
