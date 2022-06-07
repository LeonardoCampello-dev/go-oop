package main

import "fmt"

type CheckingAccount struct {
	holder  string
	agency  int
	account int
	balance float64
}

func main() {
	firstAccount := CheckingAccount{
		"Leonardo",
		111,
		1,
		999.99,
	}

	secondAccount := CheckingAccount{
		"Tester",
		222,
		2,
		100,
	}

	fmt.Println(firstAccount.transfer(100, &secondAccount))
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
	isInvalidValue := value < 0

	if isInvalidValue {
		return "Invalid value", account.balance
	}

	account.balance += value

	return "Deposit made successfully", account.balance
}

func (account *CheckingAccount) transfer(value float64, destinationAccount *CheckingAccount) (string, float64) {
	haveEnoughBalance := value > 0 && value < account.balance

	if haveEnoughBalance {
		account.balance -= value

		destinationAccount.deposit(value)

		return "Transfer performed successfully", account.balance
	}

	return "Insufficient amount for transfer", account.balance
}
