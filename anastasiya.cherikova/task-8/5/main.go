package main

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var files embed.FS // Embed all txt files

func main() {
	// Reading the list of embedded files
	dir, _ := files.ReadDir(".")

	fmt.Println("Embedded files:")
	for _, f := range dir {
		content, _ := files.ReadFile(f.Name())
		fmt.Printf("=== %s ===\n%s\n", f.Name(), content)
	}
}
