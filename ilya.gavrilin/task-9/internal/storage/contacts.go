package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"task-9/internal/model"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound          = errors.New("contact not found")
	ErrDuplicateContact  = errors.New("contact already exists")
	ErrInvalidData       = errors.New("invalid contact data")
	ErrDatabaseOperation = errors.New("database operation failed")
)

type ContactStore struct {
	db *Database
}

func NewContactStore(db *Database) *ContactStore {
	return &ContactStore{db: db}
}

func (s *ContactStore) Initialize() error {
	return s.db.RunMigrations()
}

func (s *ContactStore) GetAll() ([]model.Contact, error) {
	query := `
		SELECT id, name, phone, email, created_at, updated_at 
		FROM contacts 
		ORDER BY name
	`

	rows, err := s.db.GetConnection().Query(query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	defer rows.Close()

	var contacts []model.Contact
	for rows.Next() {
		var c model.Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Email, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
		}
		contacts = append(contacts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return contacts, nil
}

func (s *ContactStore) GetByID(id string) (model.Contact, error) {
	query := `
		SELECT id, name, phone, email, created_at, updated_at 
		FROM contacts 
		WHERE id = $1
	`

	var c model.Contact
	err := s.db.GetConnection().QueryRow(query, id).Scan(
		&c.ID, &c.Name, &c.Phone, &c.Email, &c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Contact{}, ErrNotFound
		}
		return model.Contact{}, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return c, nil
}

func (s *ContactStore) Create(contact model.Contact) (model.Contact, error) {
	if contact.Name == "" || contact.Phone == "" {
		return model.Contact{}, ErrInvalidData
	}

	if contact.ID == "" {
		contact.ID = uuid.New().String()
	}

	now := time.Now()
	contact.CreatedAt = now
	contact.UpdatedAt = now

	query := `
		INSERT INTO contacts (id, name, phone, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := s.db.GetConnection().QueryRow(
		query,
		contact.ID,
		contact.Name,
		contact.Phone,
		contact.Email,
		contact.CreatedAt,
		contact.UpdatedAt,
	).Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return model.Contact{}, ErrDuplicateContact
		}
		return model.Contact{}, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return contact, nil
}

func (s *ContactStore) Update(id string, contact model.Contact) (model.Contact, error) {
	if contact.Name == "" || contact.Phone == "" {
		return model.Contact{}, ErrInvalidData
	}

	_, err := s.GetByID(id)
	if err != nil {
		return model.Contact{}, err
	}

	contact.ID = id
	contact.UpdatedAt = time.Now()

	query := `
		UPDATE contacts
		SET name = $1, phone = $2, email = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, created_at, updated_at
	`

	err = s.db.GetConnection().QueryRow(
		query,
		contact.Name,
		contact.Phone,
		contact.Email,
		contact.UpdatedAt,
		id,
	).Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return model.Contact{}, ErrDuplicateContact
		}
		return model.Contact{}, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return contact, nil
}

func (s *ContactStore) Delete(id string) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM contacts WHERE id = $1`

	result, err := s.db.GetConnection().Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
