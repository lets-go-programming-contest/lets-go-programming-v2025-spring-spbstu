//go:build ignore

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "main.go", nil, parser.SkipObjectResolution)
	ast.Fprint(os.Stdout, fset, node, nil)
}
