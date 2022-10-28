package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

// NewAccount creates account struct
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on an account
func (a *Account /* pointer receiver */) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from an account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change owner of an account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Balance of an account
func (a Account) Balance() int {
	return a.balance
}

// Owner of an account
func (a Account) Owner() string {
	return a.owner
}

// String method
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account has ", a.balance)
}
