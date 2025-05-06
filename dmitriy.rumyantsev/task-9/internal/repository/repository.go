package repository

import (
	"context"

	"github.com/dmitriy.rumyantsev/task-9/internal/domain"
)

// ContactRepository represents an interface for working with contacts in storage
type ContactRepository interface {
	// GetAll returns all contacts
	GetAll(ctx context.Context) ([]domain.Contact, error)

	// GetByID returns a contact by ID
	GetByID(ctx context.Context, id int) (domain.Contact, error)

	// Create creates a new contact
	Create(ctx context.Context, contact domain.Contact) (domain.Contact, error)

	// Update updates an existing contact
	Update(ctx context.Context, contact domain.Contact) error

	// Delete deletes a contact by ID
	Delete(ctx context.Context, id int) error
}

// Repositories contains all repositories of the application
type Repositories struct {
	Contact ContactRepository
}

// NewRepositories creates an instance of Repositories with the specified contact repository
func NewRepositories(contactRepo ContactRepository) *Repositories {
	return &Repositories{
		Contact: contactRepo,
	}
}
