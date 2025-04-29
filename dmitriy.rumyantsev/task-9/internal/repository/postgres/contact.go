package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dmitriy.rumyantsev/task-9/internal/domain"
)

// ContactRepository implements the repository.ContactRepository interface using PostgreSQL
type ContactRepository struct {
	db *sql.DB
}

// NewContactRepository creates a new instance of ContactRepository
func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{
		db: db,
	}
}

// GetAll returns all contacts from the database
func (r *ContactRepository) GetAll(ctx context.Context) ([]domain.Contact, error) {
	query := `SELECT id, name, phone FROM contacts`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []domain.Contact
	for rows.Next() {
		var c domain.Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

// GetByID returns a contact by ID
func (r *ContactRepository) GetByID(ctx context.Context, id int) (domain.Contact, error) {
	query := `SELECT id, name, phone FROM contacts WHERE id = $1`

	var c domain.Contact
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Contact{}, sql.ErrNoRows
		}
		return domain.Contact{}, err
	}

	return c, nil
}

// Create creates a new contact
func (r *ContactRepository) Create(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	query := `INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, contact.Name, contact.Phone).Scan(&contact.ID)
	if err != nil {
		return domain.Contact{}, err
	}

	return contact, nil
}

// Update updates an existing contact
func (r *ContactRepository) Update(ctx context.Context, contact domain.Contact) error {
	query := `UPDATE contacts SET name = $1, phone = $2 WHERE id = $3`

	result, err := r.db.ExecContext(ctx, query, contact.Name, contact.Phone, contact.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete removes a contact by ID
func (r *ContactRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM contacts WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// InitSchema creates the necessary tables in the database if they don't exist
func (r *ContactRepository) InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		phone TEXT NOT NULL
	);`

	_, err := r.db.Exec(query)
	return err
}
