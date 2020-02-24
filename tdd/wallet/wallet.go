package wallet

// Wallet ...
type Wallet struct {
	balance int
}

// Deposit ...
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance ...
func (w *Wallet) Balance() int {
	return w.balance
}
