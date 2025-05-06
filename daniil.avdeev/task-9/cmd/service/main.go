package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/realFrogboy/task-9/internal/config"
	"github.com/realFrogboy/task-9/internal/db"
	"github.com/realFrogboy/task-9/internal/handler"

	"github.com/gorilla/mux"
)

func readConfig(path string) config.Config {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read config file: %s", err)
	}

	config, err := config.Parse(configFile)
	if err != nil {
		log.Fatalf("Can't parse config file: %s", err)
	}

	return config
}

func initService(dbPath string) (*db.SQLiteStorage, *mux.Router) {
	dbStorage, err := db.NewSQLiteStorage(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	contactHandler := handler.NewContactHandler(dbStorage)

	router := mux.NewRouter()

	router.HandleFunc("/contacts", contactHandler.GetAllContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", contactHandler.GetContact).Methods("GET")
	router.HandleFunc("/contacts", contactHandler.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactHandler.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contactHandler.DeleteContact).Methods("DELETE")

	return dbStorage, router
}

func main() {
	configFilePath := flag.String("config", "configs/config.yaml", "Path to the config file")
	flag.Parse()

	config := readConfig(*configFilePath)

	dbStorage, router := initService(config.DBPath)
	defer dbStorage.Close()

	log.Fatal(http.ListenAndServe(config.Port, router))
}
