package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/quaiion/go-practice/contact-manager/internal/cm"
	"github.com/quaiion/go-practice/contact-manager/internal/config"
	"github.com/quaiion/go-practice/contact-manager/internal/db"
	"github.com/quaiion/go-practice/contact-manager/internal/handler"
)

var (
        errFailedOpenBD = errors.New("failed to open db")
        errFailedInitDB = errors.New("failed to init db")
        errConfigFailed = errors.New("failed to configure")
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

        hand := handler.New(contMan)

        http.HandleFunc("/contacts", hand.HandleAllContacts)
        http.HandleFunc("/contacts/", hand.HandleContact)

        log.Print("\ncontact manager online\n")
        http.ListenAndServe(":" + configParams.ServicePort, nil)
}
