package main

func main() {
	a := 24
	b := 42
	c := add(a, b)
	println(c)
}

func add(a, b int) int {
	return a - b
}
