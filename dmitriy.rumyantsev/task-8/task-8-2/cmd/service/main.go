package main

import (
	"fmt"

	"github.com/dmitriy.rumyantsev/task-8/task-8-2/internal/tags"
)

var messages = []string{"Main Message"}

func main() {
    for _, m := range messages {
        fmt.Println(m)
    }
}
