package simulator

import (
	"fmt"
	"strings"
	"time"

	"github.com/dmitriy.rumyantsev/task-4/internal/safebank"
	"github.com/dmitriy.rumyantsev/task-4/internal/unsafebank"
)

// Simulation constants
const (
	NumAccounts       = 10   // Number of accounts in the bank
	InitialBalance    = 1000 // Initial balance for each account
	NumTransactions   = 1000 // Number of transactions to simulate
	MaxTransferAmount = 100  // Maximum amount to transfer in a transaction
)

// Provides methods to run banking simulations
type BankSimulator struct {
	SafeBank   *safebank.Bank
	UnsafeBank *unsafebank.Bank
}

// Creates a new simulator with initialized banks
func NewBankSimulator() *BankSimulator {
	return &BankSimulator{
		SafeBank:   safebank.NewBank(NumAccounts, InitialBalance),
		UnsafeBank: unsafebank.NewBank(NumAccounts, InitialBalance),
	}
}

// Runs a simulation using the safe bank implementation
func (s *BankSimulator) runSafeBankSimulation() time.Duration {
	fmt.Println("Running simulation with safe bank...")

	// Run transactions
	startTime := time.Now()
	s.SafeBank.RunSimulation(NumTransactions)
	duration := time.Since(startTime)

	// Print results
	s.printSafeBankResults(duration)

	return duration
}

// Runs a simulation using the unsafe bank implementation
func (s *BankSimulator) runUnsafeBankSimulation() time.Duration {
	fmt.Println("Running simulation with unsafe bank...")

	// Run transactions
	startTime := time.Now()
	s.UnsafeBank.RunSimulation(NumTransactions)
	duration := time.Since(startTime)

	// Print results
	s.printUnsafeBankResults(duration)

	return duration
}

// Prints the simulation results for the safe bank
func (s *BankSimulator) printSafeBankResults(duration time.Duration) {
	fmt.Printf("=== Safe Bank Simulation Results ===\n")
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("Total transactions: %d\n", s.SafeBank.GetTxCounter())
	fmt.Printf("Successful transactions: %d\n", s.SafeBank.GetSuccessTx())
	fmt.Printf("Failed transactions: %d\n", s.SafeBank.GetFailedTx())
	fmt.Printf("Initial total balance: %d\n", s.SafeBank.GetInitialBalance())
	fmt.Printf("Final total balance: %d\n", s.SafeBank.GetCurrentTotalBalance())
	fmt.Printf("Balance difference: %d\n", s.SafeBank.GetCurrentTotalBalance()-s.SafeBank.GetInitialBalance())

	// Print account balances
	fmt.Println("\nAccount balances:")
	for i, balance := range s.SafeBank.GetAllBalances() {
		fmt.Printf("Account %d: %d\n", i, balance)
	}
}

// Prints the simulation results for the unsafe bank
func (s *BankSimulator) printUnsafeBankResults(duration time.Duration) {
	fmt.Printf("=== Unsafe Bank Simulation Results ===\n")
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("Total transactions: %d\n", s.UnsafeBank.GetTxCounter())
	fmt.Printf("Successful transactions: %d\n", s.UnsafeBank.GetSuccessTx())
	fmt.Printf("Failed transactions: %d\n", s.UnsafeBank.GetFailedTx())
	fmt.Printf("Initial total balance: %d\n", s.UnsafeBank.GetInitialBalance())
	fmt.Printf("Final total balance: %d\n", s.UnsafeBank.GetCurrentTotalBalance())
	fmt.Printf("Balance difference: %d\n", s.UnsafeBank.GetCurrentTotalBalance()-s.UnsafeBank.GetInitialBalance())

	// Print account balances
	fmt.Println("\nAccount balances:")
	for i, balance := range s.UnsafeBank.GetAllBalances() {
		fmt.Printf("Account %d: %d\n", i, balance)
	}
}

// Prints a separator line
func (s *BankSimulator) printSeparator() {
	fmt.Println("\n" + strings.Repeat("-", 80) + "\n")
}

// Runs both safe and unsafe simulations
func (s *BankSimulator) RunBothSimulations() {
	fmt.Println("=== Banking System Simulator ===")

	s.runSafeBankSimulation()
	s.printSeparator()
	s.runUnsafeBankSimulation()
	s.printConclusion()
}

// Prints a conclusion comparing both simulations
func (s *BankSimulator) printConclusion() {
	fmt.Println("\n=== Simulation Conclusion ===")

	safeBalanceDiff := s.SafeBank.GetCurrentTotalBalance() - s.SafeBank.GetInitialBalance()
	unsafeBalanceDiff := s.UnsafeBank.GetCurrentTotalBalance() - s.UnsafeBank.GetInitialBalance()

	fmt.Println("Safe bank balance difference:", safeBalanceDiff)
	fmt.Println("Unsafe bank balance difference:", unsafeBalanceDiff)
}
