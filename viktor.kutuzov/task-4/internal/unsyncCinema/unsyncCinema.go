package unsyncCinema

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type UnsafeCinema struct {
	seats []string
}

func NewUnsafeCinema(totalSeats int) *UnsafeCinema {
	return &UnsafeCinema{
		seats: make([]string, totalSeats),
	}
}

func (c *UnsafeCinema) BookSeat(seat int, user string) bool {
	if c.seats[seat] == "" {
		time.Sleep(time.Microsecond * 100) // Задержка - время на бронирование
		c.seats[seat] = user
		return true
	}
	return false
}

func Demo() {
	const MaxSeats = 6
	cinema := NewUnsafeCinema(MaxSeats)
	var wg sync.WaitGroup
	users := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank"}

	baseSeed := uint64(time.Now().UnixNano())

	for i, user := range users {
		wg.Add(1)
		go func(u string, idx int) {
			defer wg.Done()

			rng := rand.New(rand.NewPCG(baseSeed+uint64(idx), uint64(time.Now().UnixNano()))) // Уникальный генератор для каждой горутины
			perm := rng.Perm(MaxSeats)                                                        // Уникальная последовательность мест для каждой горутины

			for _, seat := range perm {
				if cinema.BookSeat(seat, u) {
					fmt.Printf("%s занял место %d\n", u, seat+1)
					return
				}
			}
			fmt.Printf("%s не нашел места\n", u)
		}(user, i)
	}

	wg.Wait()

	fmt.Println("\nИтоговое распределение (с гонками):")
	for i, user := range cinema.seats {
		if user != "" {
			fmt.Printf("Место %d: %s\n", i+1, user)
		}
	}
}
