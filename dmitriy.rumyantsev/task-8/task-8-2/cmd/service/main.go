package main

import "fmt"

var messages = []string{"Hello from Main"}

func main() {
    for _, m := range messages {
        fmt.Println(m)
    }
}
