package main

func fib(n int) int {
  if (n == 0) {
    return 0
  }

  if (n == 1) {
    return 1
  }

  return fib(n - 2) + fib(n - 1)
}

func main() {
  fib(5)
}
