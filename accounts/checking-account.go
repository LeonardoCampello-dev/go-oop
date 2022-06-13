package accounts

import "github.com/LeonardoCampello-dev/go-oop/customers"

type CheckingAccount struct {
	Holder  customers.Holder
	Agency  int
	Account int
	Balance float64
}

func (account *CheckingAccount) Withdraw(value float64) (string, float64) {
	isReadyToWithdraw := value > 0 && value <= account.Balance

	if isReadyToWithdraw {
		account.Balance -= value

		return "Withdrawal successful!", account.Balance
	}

	return "Insufficient balance for withdrawal.", account.Balance
}

func (account *CheckingAccount) Deposit(value float64) (string, float64) {
	isInvalidValue := value < 0

	if isInvalidValue {
		return "Invalid value", account.Balance
	}

	account.Balance += value

	return "Deposit made successfully", account.Balance
}

func (account *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) (string, float64) {
	haveEnoughBalance := value > 0 && value < account.Balance

	if haveEnoughBalance {
		account.Balance -= value

		destinationAccount.Deposit(value)

		return "Transfer performed successfully", account.Balance
	}

	return "Insufficient amount for transfer", account.Balance
}
