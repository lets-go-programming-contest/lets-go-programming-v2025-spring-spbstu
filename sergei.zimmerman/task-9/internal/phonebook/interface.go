package phonebook

import "github.com/google/uuid"

type Contact struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
}

type Repository interface {
	GetAll() ([]Contact, error)
	GetByID(id uuid.UUID) (Contact, error)
	Create(contact Contact) error
	Update(id uuid.UUID, contact Contact) error
	Delete(id uuid.UUID) error
}
