package wallet

import (
	"errors"
	"fmt"
)

// Stringer ...
type Stringer interface {
	String() string
}

// Bitcoin ...
type Bitcoin int

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("Insufficient funds to withdraw from wallet")
	}

	w.balance -= amount
	return nil
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
