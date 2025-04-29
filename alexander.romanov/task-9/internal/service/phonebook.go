package service

import (
	"fmt"
	"github.com/alexander-romanov-edu/phonebook/internal/repository"
	"github.com/alexander-romanov-edu/phonebook/pkg/models"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^\+?[0-8][0-9]{10}`)

type PhonebookService struct {
	repo repository.PhonebookRepository
}

func NewPhonebookService(r repository.PhonebookRepository) *PhonebookService {
	return &PhonebookService{repo: r}
}

func (s *PhonebookService) Initialize() error {
	return s.repo.Initialize()
}

func (s *PhonebookService) DeleteContact(id string) error {
	return s.repo.DeleteContact(id)
}

func (s *PhonebookService) AddContact(contact models.Contact) error {
	if !phoneRegex.MatchString(contact.Phone) {
		return fmt.Errorf("invalid phone number format. Please use formats like: +71234567890")
	}
	return s.repo.AddContact(contact)
}

func (s *PhonebookService) GetContact(id string) (models.Contact, error) {
	return s.repo.GetContact(id)
}

func (s *PhonebookService) UpdateContact(contact models.Contact) error {
	if !phoneRegex.MatchString(contact.Phone) {
		return fmt.Errorf("invalid phone number format. Please use formats like: +71234567890")
	}
	return s.repo.UpdateContact(contact)
}

func (s *PhonebookService) GetAllContacts() ([]models.Contact, error) {
	return s.repo.GetAllContacts()
}
