package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-oop/accounts"
	"github.com/LeonardoCampello-dev/go-oop/customers"
)

func main() {
	leonardo := customers.Holder{
		Name:           "Leonardo",
		DocumentNumber: "780.219.590-02",
		Profession:     "Developer",
	}

	checkingAccount := accounts.CheckingAccount{
		Holder:  leonardo,
		Agency:  111,
		Account: 1,
	}

	checkingAccount.Deposit(400)

	fmt.Println(checkingAccount, checkingAccount.GetBalance())

	bianca := customers.Holder{
		Name:           "Bianca",
		DocumentNumber: "925.885.450-09",
		Profession:     "Intern",
	}

	savingsAccount := accounts.SavingsAccount{
		Holder:  bianca,
		Agency:  111,
		Account: 1,
	}

	savingsAccount.Deposit(300)
	savingsAccount.Withdraw(100)

	fmt.Println(savingsAccount.GetBalance())
}
