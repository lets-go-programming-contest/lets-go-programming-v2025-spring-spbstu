package main

import "fmt"

var (
	version = "dev"
	date    = "unknown"
)

func main() {
	fmt.Printf("Version: %s\nDate: %s\n", version, date)
}