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

	account := accounts.CheckingAccount{
		Holder:  leonardo,
		Agency:  111,
		Account: 1,
		Balance: 9999.99,
	}

	fmt.Println(account)
}
