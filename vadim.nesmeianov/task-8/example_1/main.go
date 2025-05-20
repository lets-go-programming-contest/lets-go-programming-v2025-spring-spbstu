package main

func foo(a int) int {
	b := 9 * a
	a = 2*a + b
	return a
}

func main() {
	a := 10
	_ = foo(a)
}
