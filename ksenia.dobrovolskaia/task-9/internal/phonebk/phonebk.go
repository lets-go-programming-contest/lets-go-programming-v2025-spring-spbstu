package phonebk

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Phonebook struct {
	db *sql.DB
}

func NewPhonebook(db *sql.DB) *Phonebook {
	return &Phonebook{db: db}
}

func (pb Phonebook) Initialize() error {
	query := `
  CREATE TABLE IF NOT EXISTS contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    CONSTRAINT unique_name_phone UNIQUE (name, phone)
  )`
	_, err := pb.db.Exec(query)
	return err
}

var phoneRegex = regexp.MustCompile(`^\+?[0-9]{11}`)

func (pb Phonebook) Add(contact Contact) error {
	if !phoneRegex.MatchString(contact.Phone) {
		return errors.New("phone number invalid format, need: 89141228331")
	}
	if contact.ID == "" {
		query := `
        INSERT INTO contacts (name, phone)
        VALUES ($1, $2)
        RETURNING id`

		err := pb.db.QueryRow(query, contact.Name, contact.Phone).Scan(&contact.ID)

		if err != nil {
			if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
				return errors.New("contact with this name already exists")
			}
		}
		return err
	}
	query := `
      INSERT INTO contacts (id, name, phone)
      VALUES ($1, $2, $3)`
	_, err := pb.db.Exec(query, contact.ID, contact.Name, contact.Phone)

	if err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return errors.New("contact with this name already exists")
		}
	}
	return err
}

func (pb Phonebook) Delete(id string) error {
	query := `DELETE FROM contacts WHERE id = $1`
	result, err := pb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("contact not found")
	}

	return nil
}

func (pb Phonebook) Get(id string) (Contact, error) {
	var contact Contact
	query := `SELECT id, name, phone FROM contacts WHERE id = $1`
	err := pb.db.QueryRow(query, id).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if err != nil && err == sql.ErrNoRows {
		return contact, errors.New("contact not found")
	}
	return contact, err
}

func (pb Phonebook) Update(contact Contact) error {
	if !phoneRegex.MatchString(contact.Phone) {
		return fmt.Errorf("phone number invalid format, need: 89141228331")
	}

	query := `UPDATE contacts SET name = $1, phone = $2 WHERE id = $3`
	result, err := pb.db.Exec(query, contact.Name, contact.Phone, contact.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("contact not found")
	}

	return nil
}

func (pb Phonebook) GetAll() ([]Contact, error) {
	query := `SELECT id, name, phone FROM contacts ORDER BY name`
	rows, err := pb.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
