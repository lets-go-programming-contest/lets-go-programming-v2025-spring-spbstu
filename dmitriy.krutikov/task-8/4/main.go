package main

//go:generate echo "Generating code..."
//go:generate sh -c "go version > version.txt"

import "fmt"

func main() {
	fmt.Println("Generated file with version info")
}