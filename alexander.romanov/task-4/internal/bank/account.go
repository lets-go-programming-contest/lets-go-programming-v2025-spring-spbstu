package bank

import (
  "fmt"
  "sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (ba *BankAccount) Deposit(amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.balance += amount
	fmt.Printf("Deposited $%d. New balance: $%d\n", amount, ba.balance)
}

func (ba *BankAccount) GetBalance() int {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	return ba.balance
}

