// File: main.go
package main

import (
	"fmt"
	"sync"
	"time"

	"task-4/internal/ticketsystem"
)

const (
	Seats      = 10
	Passengers = 12
)

func main() {
	demoSafeSystem()
	fmt.Println("\n=================================\n")
	demoUnsafeSystem()
}

func demoSafeSystem() {
	fmt.Println("Safe Ticket System Demo (with mutex)")
	ts := ticketsystem.NewTicketSystem(Seats)
	var wg sync.WaitGroup

	for i := 0; i < Passengers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if seat := ts.BookTicket(); seat != -1 {
				fmt.Printf("Booking %d: Seat %d booked\n", id, seat)
			} else {
				fmt.Printf("Booking %d: No seats available\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nFinal available seats: %d (expected 0)\n", ts.AvailableSeats())
}

func demoUnsafeSystem() {
	fmt.Println("Unsafe Ticket System Demo (without mutex)")
	unsafeSeats := make([]bool, Seats)
	bookingChan := make(chan int, 100) // Channel to collect bookings
	var wg sync.WaitGroup

	// Simulate concurrent bookings with random delays
	for i := 0; i < Passengers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Microsecond * time.Duration(id%10)) // Varying delays

			for seatNum := 0; seatNum < Seats; seatNum++ {
				if !unsafeSeats[seatNum] {
					// Artificial delay to create race window
					time.Sleep(time.Microsecond * 50)
					unsafeSeats[seatNum] = true
					bookingChan <- seatNum + 1
					return
				}
			}
			bookingChan <- -1
		}(i)
	}

	go func() {
		wg.Wait()
		close(bookingChan)
	}()

	// Process bookings
	bookings := make([]int, 0)
	for seat := range bookingChan {
		if seat != -1 {
			bookings = append(bookings, seat)
		}
	}

	// Calculate remaining seats
	remaining := 0
	for _, booked := range unsafeSeats {
		if !booked {
			remaining++
		}
	}
	fmt.Printf("\nFinal available seats: %d (expected 0)\n", remaining)

	// Error detection
	errors := detectErrors(unsafeSeats, bookings, Seats)
	printErrors(errors)
}

// Helper function to detect various error types
func detectErrors(seats []bool, bookings []int, totalSeats int) map[string]interface{} {
	errors := make(map[string]interface{})
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)

	// Check for duplicate bookings
	for _, seat := range bookings {
		if seen[seat] {
			duplicates[seat] = true
		} else {
			seen[seat] = true
		}
	}

	// Check seat availability consistency
	bookedCount := 0
	for _, booked := range seats {
		if booked {
			bookedCount++
		}
	}

	// Populate errors
	if len(duplicates) > 0 {
		errors["double_bookings"] = duplicates
	}
	if bookedCount != len(seen) {
		errors["count_mismatch"] = []int{bookedCount, len(seen)}
	}
	if bookedCount > totalSeats {
		errors["overbooking"] = bookedCount
	}

	return errors
}

// Helper function to print detected errors
func printErrors(errors map[string]interface{}) {
	if len(errors) == 0 {
		fmt.Println("\nNo errors detected (this is unexpected in unsafe system)")
		return
	}

	fmt.Println("\nERRORS DETECTED:")
	if dups, ok := errors["double_bookings"].(map[int]bool); ok {
		fmt.Printf(" - Double bookings: %v\n", getKeys(dups))
	}
	if mismatch, ok := errors["count_mismatch"].([]int); ok {
		fmt.Printf(" - Booked seats mismatch: system shows %d, actual unique bookings %d\n", mismatch[0], mismatch[1])
	}
	if over, ok := errors["overbooking"].(int); ok {
		fmt.Printf(" - Overbooking detected: %d bookings for %d seats\n", over, 50)
	}
}

// Helper function to get map keys
func getKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
