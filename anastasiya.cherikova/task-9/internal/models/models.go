package models

import (
	"database/sql"
	"errors"
)

// Represent a phone contact
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Create the contacts table
func InitDB(db *sql.DB) error {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS contacts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            phone TEXT NOT NULL UNIQUE
        );
    `
	if _, err := db.Exec(createTableSQL); err != nil {
		return err
	}
	return nil
}

// Retrieve all contacts
func GetAllContacts(db *sql.DB) ([]Contact, error) {
	rows, err := db.Query("SELECT id, name, phone FROM contacts ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contacts := []Contact{}
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}

// Retrieve a contact by ID
func GetContactByID(db *sql.DB, id int) (Contact, error) {
	var c Contact
	row := db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = ?", id)
	err := row.Scan(&c.ID, &c.Name, &c.Phone)
	if errors.Is(err, sql.ErrNoRows) {
		return c, errors.New("contact not found")
	}
	return c, err
}

// Add a new contact
func CreateContact(db *sql.DB, c Contact) (int64, error) {
	res, err := db.Exec("INSERT INTO contacts (name, phone) VALUES (?, ?)", c.Name, c.Phone)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Update an existing contact
func UpdateContact(db *sql.DB, id int, c Contact) error {
	_, err := db.Exec("UPDATE contacts SET name = ?, phone = ? WHERE id = ?", c.Name, c.Phone, id)
	return err
}

// Remove a contact
func DeleteContact(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id = ?", id)
	return err
}
