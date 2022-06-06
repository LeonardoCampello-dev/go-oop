package main

import "fmt"

type CheckingAccount struct {
	holder  string
	agency  int
	account int
	balance float64
}

func main() {
	account := CheckingAccount{
		"Leonardo",
		111,
		1,
		999.99,
	}

	fmt.Println(account)
}
