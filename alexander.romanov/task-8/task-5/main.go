package main

import (
    _ "embed"
    "fmt"
)

//go:embed file.txt
var content string

func main() {
    fmt.Println("Embedded content:", content)
}
