package service

import (
	"context"

	"github.com/dmitriy.rumyantsev/task-9/internal/domain"
	"github.com/dmitriy.rumyantsev/task-9/internal/repository"
)

// ContactService represents an interface for business logic of contact management
type ContactService interface {
	// GetAll returns all contacts
	GetAll(ctx context.Context) ([]domain.Contact, error)

	// GetByID returns a contact by ID
	GetByID(ctx context.Context, id int) (domain.Contact, error)

	// Create creates a new contact
	Create(ctx context.Context, contact domain.Contact) (domain.Contact, error)

	// Update updates an existing contact
	Update(ctx context.Context, id int, contact domain.Contact) (domain.Contact, error)

	// Delete removes a contact by ID
	Delete(ctx context.Context, id int) error
}

// Services contains all services of the application
type Services struct {
	Contact ContactService
}

// Dependencies contains all external dependencies for services
type Dependencies struct {
	Repos *repository.Repositories
}

// NewServices creates an instance of Services with configured services
func NewServices(deps Dependencies) *Services {
	contactService := NewContactService(deps.Repos.Contact)

	return &Services{
		Contact: contactService,
	}
}
