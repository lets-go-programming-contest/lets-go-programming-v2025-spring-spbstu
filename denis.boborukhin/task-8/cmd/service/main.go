package main

import (
	"fmt"

	"github.com/denisboborukhin/compiling/internal/data"
	"github.com/denisboborukhin/compiling/internal/tags"
)

var Version string = "default"

//go:generate go run ../../internal/generate/generate.go
func main() {
	fmt.Println("Version:", Version)
	fmt.Println(string(data.Hello), string(data.World))
	tags.PrintTags()
}
