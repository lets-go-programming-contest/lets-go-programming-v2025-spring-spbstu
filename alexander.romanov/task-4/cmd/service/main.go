package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
  "task-4/internal/bank"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	account := bank.BankAccount{}
	numCustomers := 5

	var wg sync.WaitGroup
	wg.Add(numCustomers)

	// Simulate multiple customers depositing money concurrently
	for i := 1; i <= numCustomers; i++ {
		go func(customerID int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				amount := customerID * 100

				delay := time.Duration(rand.Intn(401)+100) * time.Millisecond
				time.Sleep(delay)

				account.Deposit(amount)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Final bank account balance: $%d\n", account.GetBalance())
}
