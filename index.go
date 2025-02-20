package bank

import (
	"errors"
)

type Account struct {
	AccountNumber string
	Owner         string
	Balance       float64
}

// Deposit adds the given amount to the account balance.
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

// Withdraw subtracts the given amount from the account balance if sufficient funds are available.
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// Transfer sends the specified amount to another account if sufficient funds are available.
func (a *Account) Transfer(amount float64, recipient *Account) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	recipient.Deposit(amount)
	return nil
}

// GetBalance returns the current balance of the account.
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// NewAccount creates a new bank account with the given details.
func NewAccount(accountNumber string, owner string, initialDeposit float64) (*Account, error) {
	if initialDeposit < 0 {
		return nil, errors.New("initial deposit must not be negative")
	}

	return &Account{
		AccountNumber: accountNumber,
		Owner:         owner,
		Balance:       initialDeposit,
	}, nil
}
