package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var InsufficientFundsError = errors.New("tried to withdraw insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
