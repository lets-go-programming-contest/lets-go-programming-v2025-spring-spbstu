package main

import (
	"errors"

	grpcserver "github.com/quaiion/go-practice/grpc-contact-manager/gen/proto/contact_manager/v1"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	config "github.com/quaiion/go-practice/grpc-contact-manager/internal/configs/client"
)

var (
        errExecFailed   = errors.New("failed to execute")
        errConfigFailed = errors.New("failed to configure")
)

func main() {
        configParams, err := config.GetConfigParams()
        if err != nil {
                panic(errors.Join(errConfigFailed, err))
        }

        address := `localhost:` + configParams.ServicePort
	cmd := grpcserver.ContactManagerServiceClientCommand(client.WithServerAddr(address))

        err = cmd.Execute()
	if err != nil {
		panic(errors.Join(errExecFailed, err))
	}
}
