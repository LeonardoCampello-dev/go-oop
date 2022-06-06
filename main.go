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

	fmt.Println(account.withdraw(2000))
}

func (account *CheckingAccount) withdraw(value float64) string {
	isReadyToWithdraw := value > 0 && value <= account.balance

	if isReadyToWithdraw {
		account.balance -= value

		return "Withdrawal successful!"
	}

	return "Insufficient balance for withdrawal."
}
