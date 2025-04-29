package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/dmitriy.rumyantsev/task-9/internal/domain"
	"github.com/dmitriy.rumyantsev/task-9/internal/repository"
	"github.com/dmitriy.rumyantsev/task-9/pkg/httputil"
)

// contactService implements the ContactService interface
type contactService struct {
	repo repository.ContactRepository
}

// NewContactService creates a new instance of ContactService
func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{
		repo: repo,
	}
}

// validateName checks if the contact name is valid
func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return httputil.NewBadRequestError(
			"Invalid name",
			"Name cannot be empty",
		)
	}
	return nil
}

// validatePhone checks if the phone number is valid
func validatePhone(phone string) error {
	// Phone number validation using regular expression
	// Supported formats: +7XXXXXXXXXX, 8XXXXXXXXXX, XXXXXXXXXX
	phoneRegex := regexp.MustCompile(`^(\+7|8)?[0-9]{10}$`)
	if !phoneRegex.MatchString(phone) {
		return httputil.NewBadRequestError(
			"Invalid phone number",
			"Phone must be in format: +7XXXXXXXXXX, 8XXXXXXXXXX or XXXXXXXXXX",
		)
	}
	return nil
}

// validateContact checks the validity of a contact
func (s *contactService) validateContact(contact domain.Contact) error {
	if err := validateName(contact.Name); err != nil {
		return err
	}

	if err := validatePhone(contact.Phone); err != nil {
		return err
	}

	return nil
}

// GetAll returns all contacts
func (s *contactService) GetAll(ctx context.Context) ([]domain.Contact, error) {
	contacts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting contacts: %w", err)
	}

	return contacts, nil
}

// GetByID returns a contact by ID
func (s *contactService) GetByID(ctx context.Context, id int) (domain.Contact, error) {
	contact, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Contact{}, httputil.NewNotFoundError(
				"Contact not found",
				fmt.Sprintf("No contact found with ID: %d", id),
			)
		}
		return domain.Contact{}, fmt.Errorf("error getting contact by ID: %w", err)
	}

	return contact, nil
}

// Create creates a new contact
func (s *contactService) Create(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	if err := s.validateContact(contact); err != nil {
		return domain.Contact{}, err
	}

	newContact, err := s.repo.Create(ctx, contact)
	if err != nil {
		return domain.Contact{}, fmt.Errorf("error creating contact: %w", err)
	}

	return newContact, nil
}

// Update updates an existing contact
func (s *contactService) Update(ctx context.Context, id int, contact domain.Contact) (domain.Contact, error) {
	// First verify the contact exists
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Contact{}, err
	}

	// Then validate the new data
	if err := s.validateContact(contact); err != nil {
		return domain.Contact{}, err
	}

	// Set ID from path to the model
	contact.ID = id

	// Update the contact
	err = s.repo.Update(ctx, contact)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Contact{}, httputil.NewNotFoundError(
				"Contact not found",
				fmt.Sprintf("No contact found with ID: %d", id),
			)
		}
		return domain.Contact{}, fmt.Errorf("error updating contact: %w", err)
	}

	return contact, nil
}

// Delete removes a contact by ID
func (s *contactService) Delete(ctx context.Context, id int) error {
	// First verify the contact exists
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return httputil.NewNotFoundError(
				"Contact not found",
				fmt.Sprintf("No contact found with ID: %d", id),
			)
		}
		return fmt.Errorf("error deleting contact: %w", err)
	}

	return nil
}
