// File: ticketsystem.go
package ticketsystem

import "sync"

// TicketSystem represents a thread-safe ticket booking system
type TicketSystem struct {
	mu         sync.Mutex
	seats      []bool // Tracks booking status for each seat
	totalSeats int
}

// NewTicketSystem creates a new ticket system with given number of seats
func NewTicketSystem(totalSeats int) *TicketSystem {
	return &TicketSystem{
		seats:      make([]bool, totalSeats),
		totalSeats: totalSeats,
	}
}

// BookTicket reserves the first available seat safely
func (ts *TicketSystem) BookTicket() int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for seatNumber, booked := range ts.seats {
		if !booked {
			ts.seats[seatNumber] = true
			return seatNumber + 1 // Return 1-based seat number
		}
	}
	return -1 // No seats available
}

// AvailableSeats returns current number of available seats
func (ts *TicketSystem) AvailableSeats() int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	count := 0
	for _, booked := range ts.seats {
		if !booked {
			count++
		}
	}
	return count
}
