package wallet

import "fmt"

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

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
