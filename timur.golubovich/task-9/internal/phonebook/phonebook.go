package phonebook

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Contact struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
}

type IPhonebook interface {
	GetAll() ([]Contact, error)
	GetByID(id uuid.UUID) (Contact, error)
	Create(contact Contact) error
	Update(id uuid.UUID, contact Contact) error
	Delete(id uuid.UUID) error
}

type Phonebook struct {
	db *sql.DB
}

func New(dbPath string) (*Phonebook, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	sqlStmt := `CREATE TABLE IF NOT EXISTS contacts (
		id TEXT PRIMARY KEY,
		name TEXT,
		phone TEXT
	);`
	_, err = db.Exec(sqlStmt)
	return &Phonebook{db}, err
}

func (pb *Phonebook) GetAll() ([]Contact, error) {
	rows, err := pb.db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	contacts := []Contact{}
	for rows.Next() {
		var c Contact
		var id string
		if err := rows.Scan(&id, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		c.ID, _ = uuid.Parse(id)
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func (pb *Phonebook) GetByID(id uuid.UUID) (Contact, error) {
	var c Contact
	row := pb.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = ?", id.String())
	var sid string
	if err := row.Scan(&sid, &c.Name, &c.Phone); err != nil {
		return c, err
	}
	c.ID = id
	return c, nil
}

func (pb *Phonebook) Create(c Contact) error {
	_, err := pb.db.Exec("INSERT INTO contacts (id, name, phone) VALUES (?, ?, ?)", c.ID.String(), c.Name, c.Phone)
	return err
}

func (pb *Phonebook) Update(id uuid.UUID, c Contact) error {
	_, err := pb.db.Exec("UPDATE contacts SET name = ?, phone = ? WHERE id = ?", c.Name, c.Phone, id.String())
	return err
}

func (pb *Phonebook) Delete(id uuid.UUID) error {
	_, err := pb.db.Exec("DELETE FROM contacts WHERE id = ?", id.String())
	return err
}
