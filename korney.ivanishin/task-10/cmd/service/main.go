package main

import (
	"errors"

	"net"
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	grpcserver "github.com/quaiion/go-practice/grpc-contact-manager/gen/proto/contact_manager/v1"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/cm"
	config "github.com/quaiion/go-practice/grpc-contact-manager/internal/configs/service"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/db"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/handlers/grpchandler"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/handlers/resthandler"
	"google.golang.org/grpc"
)

var (
        errFailedOpenBD      = errors.New("failed to open db")
        errFailedInitDB      = errors.New("failed to init db")
        errFailedCreateValid = errors.New("failed to create validator")
        errFailedToListen    = errors.New("failed to listen")
        errFailedToServeGrpc = errors.New("failed to serve gRPC")
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

        go startRest(contMan, configParams)
        startGrpc(contMan, validator, configParams)
}

func startRest(contMan *cm.ContMan, configParams config.ConfigParams) {
        restHand := resthandler.New(contMan)

        http.HandleFunc("/contacts", restHand.HandleAllContacts)
        http.HandleFunc("/contacts/", restHand.HandleContact)

        restAddress := `localhost:` + configParams.RestServicePort
        http.ListenAndServe(restAddress, nil)
}

func startGrpc(contMan *cm.ContMan, validator protovalidate.Validator, configParams config.ConfigParams) {
        grpcHand := grpchandler.New(contMan, validator)

        grpcServer := grpc.NewServer()
        grpcserver.RegisterContactManagerServiceServer(grpcServer, grpcHand)

        grpcAddress := `localhost:` + configParams.GrpcServicePort
        listener, err := net.Listen("tcp", grpcAddress)
        if err != nil {
                panic(errors.Join(errFailedToListen, err))
        }

        err = grpcServer.Serve(listener)
        if err != nil {
                panic(errors.Join(errFailedToServeGrpc, err))
        }
}
