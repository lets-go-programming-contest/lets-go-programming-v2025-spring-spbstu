package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func connectDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "user=viktor dbname=contacts password=12345 sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// Get all contacts
func getContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, phone FROM contacts")
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to fetch contacts")
			return
		}
		defer rows.Close()

		var contacts []Contact
		for rows.Next() {
			var c Contact
			if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
				respondWithError(w, http.StatusInternalServerError, "Failed to scan contact")
				return
			}
			contacts = append(contacts, c)
		}
		respondWithJSON(w, http.StatusOK, contacts)
	}
}

// Get contact by ID
func getContactByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		var c Contact
		row := db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id)
		if err := row.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			if err == sql.ErrNoRows {
				respondWithError(w, http.StatusNotFound, "Contact not found")
				return
			}
			respondWithError(w, http.StatusInternalServerError, "Failed to fetch contact")
			return
		}
		respondWithJSON(w, http.StatusOK, c)
	}
}

// Create contact
func createContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c Contact
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if !isValidContact(c) {
			respondWithError(w, http.StatusBadRequest, "Name and valid phone are required")
			return
		}

		err := db.QueryRow("INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id",
			c.Name, c.Phone).Scan(&c.ID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to create contact")
			return
		}
		respondWithJSON(w, http.StatusCreated, c)
	}
}

// Update contact
func updateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		var c Contact
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if !isValidContact(c) {
			respondWithError(w, http.StatusBadRequest, "Name and valid phone are required")
			return
		}

		result, err := db.Exec("UPDATE contacts SET name=$1, phone=$2 WHERE id=$3",
			c.Name, c.Phone, id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to update contact")
			return
		}

		if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
			respondWithError(w, http.StatusNotFound, "Contact not found")
			return
		}
		c.ID = id
		respondWithJSON(w, http.StatusOK, c)
	}
}

// Delete contact
func deleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		result, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete contact")
			return
		}

		if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
			respondWithError(w, http.StatusNotFound, "Contact not found")
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Contact deleted"})
	}
}

func isValidContact(c Contact) bool {
	if c.Name == "" || c.Phone == "" {
		return false
	}
	return isValidPhone(c.Phone)
}

func isValidPhone(phone string) bool {
	matched, _ := regexp.MatchString(`^\+?[0-9]{10,15}$`, phone)
	return matched
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func main() {
	db := connectDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/contacts", getContacts(db)).Methods("GET")
	router.HandleFunc("/contacts/{id}", getContactByID(db)).Methods("GET")
	router.HandleFunc("/contacts", createContact(db)).Methods("POST")
	router.HandleFunc("/contacts/{id}", updateContact(db)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", deleteContact(db)).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
