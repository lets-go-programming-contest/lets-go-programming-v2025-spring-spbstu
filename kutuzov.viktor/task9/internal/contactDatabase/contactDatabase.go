package contactdatabase

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ConnectDB() *sql.DB {
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

func GetContacts(db *sql.DB) http.HandlerFunc {
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

func GetContactByID(db *sql.DB) http.HandlerFunc {
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

func CreateContact(db *sql.DB) http.HandlerFunc {
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
func UpdateContact(db *sql.DB) http.HandlerFunc {
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
func DeleteContact(db *sql.DB) http.HandlerFunc {
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
