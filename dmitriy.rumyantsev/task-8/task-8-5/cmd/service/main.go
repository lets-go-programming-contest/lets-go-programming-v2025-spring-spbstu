package main

import (
    _ "embed"
    "fmt"
)

//go:embed data.txt
var data string

func main() {
    fmt.Println(data)
}
