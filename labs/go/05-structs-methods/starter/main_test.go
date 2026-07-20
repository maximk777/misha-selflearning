package starter

import "testing"

func TestDepositUsesPointerReceiver(t *testing.T) {
	account := Account{Owner: "Миша", Balance: 10}
	account.Deposit(5)
	if account.Balance != 15 {
		t.Fatalf("balance=%d, want 15", account.Balance)
	}
}

func TestLabelDoesNotMutateAccount(t *testing.T) {
	account := Account{Owner: "Миша", Balance: 10}
	if got := account.Label(); got != "Миша: 10" {
		t.Fatalf("Label()=%q", got)
	}
}
