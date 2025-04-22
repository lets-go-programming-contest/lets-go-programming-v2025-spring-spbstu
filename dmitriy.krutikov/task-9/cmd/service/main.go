package main

import (
	"log"
	"net/http"
	"phonebook/internal/contact"
	"phonebook/internal/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	contactRepo := contact.NewPostgresRepository(db)
	contactService := contact.NewService(contactRepo)
	contactHandler := contact.NewHandler(contactService)

	http.HandleFunc("/contacts", contactHandler.HandleContacts)
	http.HandleFunc("/contacts/", contactHandler.HandleSingleContact)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}