package main

import (
	_ "embed"
	"log"
)

//go:embed README.md
var readme string

func main() {
	log.Print(readme)
}
