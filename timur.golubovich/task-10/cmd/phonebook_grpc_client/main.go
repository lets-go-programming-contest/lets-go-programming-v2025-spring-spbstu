package main

import (
	"fmt"
	"os"

	v1 "task-10/gen/proto/phonebook/v1"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
)

func main() {
	cmd := v1.PhonebookServiceClientCommand(
		client.WithServerAddr("localhost:50051"),
	)
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
