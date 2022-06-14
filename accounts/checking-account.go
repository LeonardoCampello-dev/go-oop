package accounts

import "github.com/LeonardoCampello-dev/go-oop/customers"

type CheckingAccount struct {
	Holder  customers.Holder
	Agency  int
	Account int
	balance float64
}

func (account *CheckingAccount) Withdraw(value float64) (string, float64) {
	isReadyToWithdraw := value > 0 && value <= account.balance

	if isReadyToWithdraw {
		account.balance -= value

		return "Withdrawal successful!", account.balance
	}

	return "Insufficient balance for withdrawal.", account.balance
}

func (account *CheckingAccount) Deposit(value float64) (string, float64) {
	isInvalidValue := value < 0

	if isInvalidValue {
		return "Invalid value", account.balance
	}

	account.balance += value

	return "Deposit made successfully", account.balance
}

func (account *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) (string, float64) {
	haveEnoughbalance := value > 0 && value < account.balance

	if haveEnoughbalance {
		account.balance -= value

		destinationAccount.Deposit(value)

		return "Transfer performed successfully", account.balance
	}

	return "Insufficient amount for transfer", account.balance
}

func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}
