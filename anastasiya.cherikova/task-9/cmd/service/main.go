package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"task-9/internal/handlers"
	"task-9/internal/models"
)

func main() {
	// Initialize database
	db, err := sql.Open("sqlite3", "./contacts.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Create tables
	if err := models.InitDB(db); err != nil {
		log.Fatal("Database initialization failed:", err)
	}

	// Setup router
	router := mux.NewRouter()

	// Register handlers
	router.HandleFunc("/contacts", handlers.GetContacts(db)).Methods("GET")
	router.HandleFunc("/contacts/{id}", handlers.GetContact(db)).Methods("GET")
	router.HandleFunc("/contacts", handlers.CreateContact(db)).Methods("POST")
	router.HandleFunc("/contacts/{id}", handlers.UpdateContact(db)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", handlers.DeleteContact(db)).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
