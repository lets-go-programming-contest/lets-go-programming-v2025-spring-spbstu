package main

func main() {
	primes := sieve(50)
	printPrimes(primes)
}

// Sieve of Eratosthenes with manual root calculation
func sieve(n int) []int {
	sieve := make([]bool, n)

	// Replacing math.Sqrt via successive multiplication
	max := 2
	for max*max < n {
		max++
	}

	for i := 2; i < max; i++ {
		if !sieve[i] {
			for j := i * i; j < n; j += i {
				sieve[j] = true
			}
		}
	}

	primes := make([]int, 0)
	for i := 2; i < n; i++ {
		if !sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Output via built-in println
func printPrimes(primes []int) {
	for _, p := range primes {
		println(p)
	}
}
