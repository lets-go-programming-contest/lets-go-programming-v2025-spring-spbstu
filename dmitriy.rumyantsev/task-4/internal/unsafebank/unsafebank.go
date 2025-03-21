package unsafebank

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Represents a bank account without thread-safety protection
type Account struct {
	ID      int
	Balance int
}

// Implements a bank without thread-safety
type Bank struct {
	Accounts       []*Account
	InitialBalance int64          // Initial total balance across all accounts
	TxCounter      int64          // Transaction counter (using atomic for counters only)
	SuccessTx      int64          // Successful transactions counter
	FailedTx       int64          // Failed transactions counter
	wg             sync.WaitGroup // WaitGroup for waiting for all transactions to complete
}

// Simulation constants
const (
	TransactionInterval = 1     // Time interval between transactions (ms)
	ShowDetailedLogs    = false // Whether to show detailed transaction logs
)

// Creates a new unsafe bank
func NewBank(numAccounts int, initialBalance int) *Bank {
	accounts := make([]*Account, numAccounts)
	var totalBalance int64

	for i := 0; i < numAccounts; i++ {
		accounts[i] = &Account{
			ID:      i,
			Balance: initialBalance,
		}
		totalBalance += int64(initialBalance)
	}

	return &Bank{
		Accounts:       accounts,
		InitialBalance: totalBalance,
	}
}

// Returns the total number of transactions
func (b *Bank) GetTxCounter() int64 {
	return atomic.LoadInt64(&b.TxCounter)
}

// Returns the number of successful transactions
func (b *Bank) GetSuccessTx() int64 {
	return atomic.LoadInt64(&b.SuccessTx)
}

// Returns the number of failed transactions
func (b *Bank) GetFailedTx() int64 {
	return atomic.LoadInt64(&b.FailedTx)
}

// Returns the initial total balance
func (b *Bank) GetInitialBalance() int64 {
	return b.InitialBalance
}

// Calculates and returns the current total balance across all accounts
func (b *Bank) GetCurrentTotalBalance() int64 {
	var totalBalance int64
	for _, acc := range b.Accounts {
		totalBalance += int64(acc.Balance) // No thread-safety here!
	}
	return totalBalance
}

// Returns a slice containing the balance of each account
func (b *Bank) GetAllBalances() []int {
	balances := make([]int, len(b.Accounts))
	for i, acc := range b.Accounts {
		balances[i] = acc.Balance // No thread-safety here!
	}
	return balances
}

// Unsafely transfers funds between accounts
func (b *Bank) transfer(fromID, toID, amount int) bool {
	// Increment transaction counter (using atomic for demonstration)
	atomic.AddInt64(&b.TxCounter, 1)

	// Parameter validation
	if fromID == toID || amount <= 0 || fromID >= len(b.Accounts) || toID >= len(b.Accounts) {
		atomic.AddInt64(&b.FailedTx, 1)
		return false
	}

	// Get accounts
	fromAccount := b.Accounts[fromID]
	toAccount := b.Accounts[toID]

	// UNSAFE: Check if enough funds available
	// Race condition: another goroutine might modify the balance between the check and the actual transfer
	if fromAccount.Balance < amount {
		atomic.AddInt64(&b.FailedTx, 1)
		if ShowDetailedLogs {
			fmt.Printf("Insufficient funds in account %d: %d < %d\n", fromID, fromAccount.Balance, amount)
		}
		return false
	}

	// UNSAFE: Perform transfer without proper synchronization
	// Race condition: multiple goroutines might modify the balance concurrently
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	// Increment successful transactions counter
	atomic.AddInt64(&b.SuccessTx, 1)

	if ShowDetailedLogs {
		fmt.Printf("Transfer %d from account %d to account %d: successful\n", amount, fromID, toID)
	}
	return true
}

// Executes a random transaction in the bank
func (b *Bank) runTransaction() {
	defer b.wg.Done()

	// Generate random transaction parameters
	numAccounts := len(b.Accounts)
	fromID := rand.Intn(numAccounts)
	toID := rand.Intn(numAccounts)
	for toID == fromID {
		toID = rand.Intn(numAccounts)
	}
	amount := rand.Intn(100) + 1 // Random amount between 1 and 100

	// Execute transaction
	b.transfer(fromID, toID, amount)

	// Small delay to simulate a real environment
	time.Sleep(time.Millisecond * time.Duration(TransactionInterval))
}

// Runs a simulation
func (b *Bank) RunSimulation(numTransactions int) {
	fmt.Printf("Starting simulation with %d transactions (unsafe version)...\n", numTransactions)

	// Start transactions in separate goroutines
	b.wg.Add(numTransactions)

	for i := 0; i < numTransactions; i++ {
		go b.runTransaction()
	}

	// Wait for all transactions to complete
	b.wg.Wait()
}
