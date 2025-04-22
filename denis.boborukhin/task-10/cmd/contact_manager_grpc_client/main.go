package main

import (
	"os"

	v1 "github.com/denisboborukhin/contact_manager/gen/proto/contact_manager/v1"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
)

func main() {
	cmd := v1.ContactManagerServiceClientCommand(
		client.WithServerAddr("localhost:50051"),
	)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
