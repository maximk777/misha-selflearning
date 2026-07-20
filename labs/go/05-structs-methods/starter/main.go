package starter

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func (account *Account) Deposit(amount int) {
	account.Balance += amount
}

func (account Account) Label() string {
	return fmt.Sprintf("%s: %d", account.Owner, account.Balance)
}
