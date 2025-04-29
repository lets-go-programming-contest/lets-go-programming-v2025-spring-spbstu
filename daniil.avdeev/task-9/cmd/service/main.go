package main

import (
	"log"
	"net/http"

	"github.com/realFrogboy/task-9/internal/handler"
	"github.com/realFrogboy/task-9/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	dbStorage, err := db.NewSQLiteStorage("phonebook.db")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer dbStorage.Close()

	contactHandler := handler.NewContactHandler(dbStorage)

	r := mux.NewRouter()

	r.HandleFunc("/contacts", contactHandler.GetAllContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", contactHandler.GetContact).Methods("GET")
	r.HandleFunc("/contacts", contactHandler.CreateContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", contactHandler.UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", contactHandler.DeleteContact).Methods("DELETE")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
