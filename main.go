package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-oop/accounts"
)

func main() {
	firstAccount := accounts.CheckingAccount{
		Holder:  "Leonardo",
		Agency:  111,
		Account: 1,
		Balance: 9999.99,
	}

	fmt.Println(firstAccount)
}
