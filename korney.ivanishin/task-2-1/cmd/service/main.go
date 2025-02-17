package main

import (
        "fmt"
        "internal/stream_utils"
        "internal/process"
)

func main() {
        err := process.ProcessOffice()
        if err != nil {
                stream_utils.FlushInput()
                fmt.Println(err)
        }
}
