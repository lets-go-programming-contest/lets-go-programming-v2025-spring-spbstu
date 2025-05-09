package main

import (
	"errors"

	// "net/http"
	"net"

	"github.com/bufbuild/protovalidate-go"
	grpcserver "github.com/quaiion/go-practice/grpc-contact-manager/gen/proto/contact_manager/v1"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/cm"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/config"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/db"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/handler"
	"google.golang.org/grpc"
)

var (
        errFailedOpenBD      = errors.New("failed to open db")
        errFailedInitDB      = errors.New("failed to init db")
        errFailedCreateValid = errors.New("failed to create validator")
        errFailedToListen    = errors.New("failed to listen")
        errFailedToServe     = errors.New("failed to serve")
        errConfigFailed      = errors.New("failed to configure")
)

func main() {
        configParams, err := config.GetConfigParams()
        if err != nil {
                panic(errors.Join(errConfigFailed, err))
        }

        database := db.New()
        err = database.Open(configParams.DBPort, configParams.DBPswd)
        if err != nil {
                panic(errors.Join(errFailedOpenBD, err))
        }
        defer database.Close()

        contMan := cm.New(database.Postgres)
        err = contMan.Init()
        if err != nil {
                panic(errors.Join(errFailedInitDB, err))
        }

        validator, err := protovalidate.New()
	if err != nil {
		panic(errors.Join(errFailedCreateValid, err))
	}

        hand := handler.New(contMan, validator)

        grpcServer := grpc.NewServer()
        grpcserver.RegisterContactManagerServiceServer(grpcServer, hand)

        address := `localhost:` + configParams.ServicePort
        listener, err := net.Listen("tcp", address)
        if err != nil {
                panic(errors.Join(errFailedToListen, err))
        }

        err = grpcServer.Serve(listener)
        if err != nil {
                panic(errors.Join(errFailedToServe, err))
        }

        // http.HandleFunc("/contacts", hand.HandleAllContacts)
        // http.HandleFunc("/contacts/", hand.HandleContact)

        // http.ListenAndServe(":" + configParams.ServicePort, nil)
}
