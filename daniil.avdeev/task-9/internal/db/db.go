package db

import (
	"database/sql"

	"github.com/realFrogboy/task-9/internal/contacts"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS contacts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			phone TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}

func (s *SQLiteStorage) GetAllContacts() ([]contacts.DBContact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allContacts []contacts.DBContact
	for rows.Next() {
		var c contacts.DBContact
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		allContacts = append(allContacts, c)
	}

	return allContacts, nil
}

func (s *SQLiteStorage) GetContactByID(id int) (*contacts.DBContact, error) {
	var c contacts.DBContact
	err := s.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func (s *SQLiteStorage) CreateContact(contact contacts.Contact) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO contacts (name, phone) VALUES (?, ?)",
		contact.Name, contact.Phone,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SQLiteStorage) UpdateContact(id int, contact contacts.Contact) error {
	_, err := s.db.Exec(
		"UPDATE contacts SET name = ?, phone = ? WHERE id = ?",
		contact.Name, contact.Phone, id,
	)
	return err
}

func (s *SQLiteStorage) DeleteContact(id int) error {
	_, err := s.db.Exec("DELETE FROM contacts WHERE id = ?", id)
	return err
}
