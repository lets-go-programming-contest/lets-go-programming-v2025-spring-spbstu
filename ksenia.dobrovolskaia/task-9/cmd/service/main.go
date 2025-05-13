package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	_ "github.com/lib/pq"

	"github.com/kseniadobrovolskaia/task-9/internal/config"
	"github.com/kseniadobrovolskaia/task-9/internal/database"
	"github.com/kseniadobrovolskaia/task-9/internal/handler"
	"github.com/kseniadobrovolskaia/task-9/internal/phonebk"
)

var configPath = flag.String("config", "", "Path to config file .yaml")

func main() {
	flag.Parse()
	if *configPath == "" {
		color.Red("--config not specified")
		os.Exit(0)
	}

	// Read config file
	cfg, err := config.ReadConfigFile(*configPath)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	bd := database.NewDataBase()
	err = bd.Open(cfg)
	if err != nil {
		color.Red("error while opening bd: " + err.Error())
		os.Exit(1)
	}
	defer bd.Close()

	phonebook := phonebk.NewPhonebook(bd.Db)
	err = phonebook.Initialize()
	if err != nil {
		color.Red("failed to init database: " + err.Error())
		os.Exit(1)
	}

	hdler := handler.NewHandler(*phonebook)

	http.HandleFunc("/contacts", hdler.HandleContacts)
	http.HandleFunc("/contacts/", hdler.HandleOneContact)

	color.Green("\nPhonebook is running!\n")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
