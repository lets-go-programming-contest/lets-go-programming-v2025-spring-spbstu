package fib

// Fibonacci returns the n-th Fibonacci number (recursive implementation).
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}