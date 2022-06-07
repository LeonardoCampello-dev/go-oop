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

	fmt.Println(account.deposit(100))
}

func (account *CheckingAccount) withdraw(value float64) (string, float64) {
	isReadyToWithdraw := value > 0 && value <= account.balance

	if isReadyToWithdraw {
		account.balance -= value

		return "Withdrawal successful!", account.balance
	}

	return "Insufficient balance for withdrawal.", account.balance
}

func (account *CheckingAccount) deposit(value float64) (string, float64) {
	if value < 0 {
		return "Invalid value", account.balance
	}

	account.balance += value

	return "Deposit made successfully", account.balance
}
