package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	// Check if a file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <go-file>")
		return
	}

	// Read the file path from command-line arguments
	filePath := os.Args[1]

	// Create a new file set
	fset := token.NewFileSet()

	// Parse the file into an AST
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}

	// Print the AST
	ast.Print(fset, node)
}
