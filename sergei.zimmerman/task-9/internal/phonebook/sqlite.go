package phonebook

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(dbPath string) (*SQLiteRepo, error) {
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

	return &SQLiteRepo{db}, err
}

func (r *SQLiteRepo) GetAll() ([]Contact, error) {
	rows, err := r.db.Query("SELECT id, name, phone FROM contacts")
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

func (r *SQLiteRepo) GetByID(id uuid.UUID) (Contact, error) {
	var c Contact

	row := r.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = ?", id.String())

	var sid string

	if err := row.Scan(&sid, &c.Name, &c.Phone); err != nil {
		return c, err
	}

	c.ID = id

	return c, nil
}

func (r *SQLiteRepo) Create(c Contact) error {
	_, err := r.db.Exec("INSERT INTO contacts (id, name, phone) VALUES (?, ?, ?)", c.ID.String(), c.Name, c.Phone)

	return err
}

func (r *SQLiteRepo) Update(id uuid.UUID, c Contact) error {
	_, err := r.db.Exec("UPDATE contacts SET name = ?, phone = ? WHERE id = ?", c.Name, c.Phone, id.String())

	return err
}

func (r *SQLiteRepo) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM contacts WHERE id = ?", id.String())

	return err
}
