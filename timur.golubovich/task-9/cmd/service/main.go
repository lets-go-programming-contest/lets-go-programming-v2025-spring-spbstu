package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"task-9/internal/config"
	handler "task-9/internal/handlers"
	"task-9/internal/phonebook"
	"task-9/internal/service"

	"github.com/gorilla/mux"
)

func app() error {
	cfg, err := config.ParseConfig("data/config.yaml")
	if err != nil {
		return fmt.Errorf("fail to parse config file: %v", err)
	}
	pb, err := phonebook.New(cfg.Path)
	if err != nil {
		return fmt.Errorf("fail to create phonebook %v", err)
	}
	svc := service.New(pb)
	h := handler.New(svc)
	r := mux.NewRouter()
	h.RegisterRoutes(r)

	log.Printf("Listening on port %v\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Port), r))
	return nil
}

func main() {
	err := app()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
