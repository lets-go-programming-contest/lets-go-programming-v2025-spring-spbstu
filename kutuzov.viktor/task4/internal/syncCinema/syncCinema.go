package syncCinema

import (
	"fmt"
	"sync"
	"time"

	"math/rand/v2"
)

type Cinema struct {
	mu    sync.RWMutex
	seats []string
	sem   chan struct{}
}

func NewCinema(capacity, maxParallel int) *Cinema {
	return &Cinema{
		seats: make([]string, capacity),
		sem:   make(chan struct{}, maxParallel),
	}
}

func (c *Cinema) BookSeat(user string, seat int) bool {
	c.sem <- struct{}{}        // ограничиваем параллельные операции
	defer func() { <-c.sem }() // освобождаем слот

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.seats[seat] == "" {
		time.Sleep(time.Microsecond * 100) // Задержка - время на бронирование
		c.seats[seat] = user
		return true
	}
	return false
}

func (c *Cinema) GetSeats() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]string, len(c.seats))
	for k, v := range c.seats {
		result[k] = v
	}
	return result
}

func Demo() {
	const MaxSeats = 6
	const MaxProc = 2

	cinema := NewCinema(MaxSeats, MaxProc)
	var wg sync.WaitGroup
	users := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Darvin"}

	baseSeed := uint64(time.Now().UnixNano())

	for i, user := range users {
		wg.Add(1)
		go func(u string, idx int) {
			defer wg.Done()

			rng := rand.New(rand.NewPCG(baseSeed+uint64(idx), uint64(time.Now().UnixNano()))) // Уникальный генератор для каждой горутины

			seats := rng.Perm(MaxSeats)

			for _, seat := range seats { //занимаем места
				if cinema.BookSeat(u, seat) {
					fmt.Printf("%s занял место %d\n", u, seat)
					return
				}
				fmt.Printf("%s не смог занять место %d\n", u, seat)
			}
		}(user, i)
	}

	wg.Wait()
	fmt.Println("\nИтоговое распределение:")
	for seat, user := range cinema.GetSeats() {
		fmt.Printf("Место %d: %s\n", seat, user)
	}
}
