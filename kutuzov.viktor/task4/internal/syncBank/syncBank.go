package syncBank

import (
	"fmt"
	"sync"
)

// Bank account with mutex
type Account struct {
	mu      sync.Mutex
	balance int
}

// Safe operation: add money
func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Safe operation: get money
func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}

// Safe operation: view balance
func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func Demo() {
	var wg sync.WaitGroup
	acc := Account{balance: 1000}
	n := 100

	wg.Add(n * 2)
	for i := 0; i < n; i++ { // parallel operations
		go func() {
			acc.Deposit(10)
			wg.Done()
		}()

		go func() {
			acc.Withdraw(10)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Synchronized:", acc.Balance())
}
