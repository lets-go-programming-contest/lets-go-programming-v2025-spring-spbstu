package contact

import (
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll() ([]Contact, error)
	GetByID(id int) (Contact, error)
	GetByName(name string) (Contact, error) 
	Create(c Contact) (Contact, error)
	Update(id int, c Contact) (Contact, error)
	DeleteByID(id int) error
	DeleteByName(name string) error 
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetByName(name string) (Contact, error) {
	var c Contact
	err := r.db.QueryRow("SELECT id, name, phone FROM contacts WHERE name = $1", name).
		Scan(&c.ID, &c.Name, &c.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return Contact{}, errors.New("contact not found")
		}
		return Contact{}, err
	}
	return c, nil
}

func (r *PostgresRepository) DeleteByName(name string) error {
	result, err := r.db.Exec("DELETE FROM contacts WHERE name = $1", name)
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

func (r *PostgresRepository) GetAll() ([]Contact, error) {
	rows, err := r.db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r *PostgresRepository) GetByID(id int) (Contact, error) {
	var c Contact
	err := r.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).
		Scan(&c.ID, &c.Name, &c.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return Contact{}, errors.New("contact not found")
		}
		return Contact{}, err
	}

	return c, nil
}

func (r *PostgresRepository) Create(c Contact) (Contact, error) {
	err := r.db.QueryRow(
		"INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id",
		c.Name, c.Phone,
	).Scan(&c.ID)

	if err != nil {
		return Contact{}, err
	}

	return c, nil
}

func (r *PostgresRepository) Update(id int, c Contact) (Contact, error) {
	result, err := r.db.Exec(
		"UPDATE contacts SET name = $1, phone = $2 WHERE id = $3",
		c.Name, c.Phone, id,
	)

	if err != nil {
		return Contact{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Contact{}, err
	}

	if rowsAffected == 0 {
		return Contact{}, errors.New("contact not found")
	}

	c.ID = id
	return c, nil
}

func (r *PostgresRepository) DeleteByID(id int) error {
	result, err := r.db.Exec("DELETE FROM contacts WHERE id = $1", id)
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