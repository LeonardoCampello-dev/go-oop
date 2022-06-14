package accounts

import "github.com/LeonardoCampello-dev/go-oop/customers"

type SavingsAccount struct {
	Holder customers.Holder
	Agency,
	Account,
	Operation int
	balance float64
}

func (account *SavingsAccount) Withdraw(value float64) (string, float64) {
	isReadyToWithdraw := value > 0 && value <= account.balance

	if isReadyToWithdraw {
		account.balance -= value

		return "Withdrawal successful!", account.balance
	}

	return "Insufficient balance for withdrawal.", account.balance
}

func (account *SavingsAccount) Deposit(value float64) (string, float64) {
	isInvalidValue := value < 0

	if isInvalidValue {
		return "Invalid value", account.balance
	}

	account.balance += value

	return "Deposit made successfully", account.balance
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
