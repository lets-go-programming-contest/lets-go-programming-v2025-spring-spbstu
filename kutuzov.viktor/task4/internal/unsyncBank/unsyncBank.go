package unsyncBank

import (
	"fmt"
	"sync"
)

func Demo() {
	var wg sync.WaitGroup
	balance := 1000
	n := 100

	wg.Add(n * 2)
	for i := 0; i < n; i++ { // parallel operations
		go func() {
			balance += 10
			wg.Done()
		}()
		go func() {
			if balance >= 10 {
				balance -= 10
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Без синхронизации:", balance)
}
