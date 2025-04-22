package example

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
  for i := 2; i <= int(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
