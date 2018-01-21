// Package account manages a bank account:
// - open an account
// - deposit or retrieve money from account
// - output the balance of an account
// - close an account
package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	sync.Mutex
	// we use mutex to allow multiple accounts
	// to be opened at the same time
	// without mixing up the values

	// embed of sync.Mutex allows for a.Lock()
	// as seen in http://exercism.io/codegoalie's solution +1

}

// Open creates and initializes a bank account
func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	a := Account{balance: initalDeposit, closed: false}
	return &a

}

// Close closes a bank account: empties it, prevents further
func (a *Account) Close() (int64, bool) {

	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	payout := a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance retrieve the amount available on a bank account
func (a *Account) Balance() (int64, bool) {

	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit allows for adding and substracting amounts from a bank account
func (a *Account) Deposit(amount int64) (int64, bool) {

	a.Lock()
	defer a.Unlock()

	if a.closed || a.balance < -amount {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
