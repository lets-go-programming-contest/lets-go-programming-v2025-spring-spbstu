package main

import (
        "fmt"
        "github.com/conditioners/internal/streamUtils"
        "github.com/conditioners/internal/process"
)

func main() {
        err := process.ProcessOffice()
        if err != nil {
                streamUtils.FlushInput()
                fmt.Println(err)
        }
}
