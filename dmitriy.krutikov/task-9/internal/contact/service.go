package contact

import (
	"errors"
	"regexp"
)

type Service interface {
	GetAllContacts() ([]Contact, error)
	GetContactByID(id int) (Contact, error)
	GetContactByName(name string) (Contact, error) 
	CreateContact(c Contact) (Contact, error)
	UpdateContact(id int, c Contact) (Contact, error)
	DeleteContactByID(id int) error
	DeleteContactByName(name string) error 
}

type contactService struct {
	repo Repository
}

var phoneRegex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

func isValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func isValidName(name string) bool {
	return len(name) >= 2 && len(name) <= 50
}

func NewService(repo Repository) Service {
	return &contactService{repo: repo}
}

func (s *contactService) GetAllContacts() ([]Contact, error) {
	return s.repo.GetAll()
}

func (s *contactService) GetContactByName(name string) (Contact, error) {
	return s.repo.GetByName(name)
}

func (s *contactService) GetContactByID(id int) (Contact, error) {
	return s.repo.GetByID(id)
}


func (s *contactService) CreateContact(c Contact) (Contact, error) {
	if !isValidName(c.Name) {
		return Contact{}, errors.New("invalid name")
	}

	if !isValidPhone(c.Phone) {
		return Contact{}, errors.New("invalid phone format")
	}

	if c.Name == "Anton" {
		return Contact{}, errors.New("name cannot be Anton")
	}

	return s.repo.Create(c)
}

func (s *contactService) UpdateContact(id int, c Contact) (Contact, error) {
	if !isValidName(c.Name) {
		return Contact{}, errors.New("invalid name")
	}

	if !isValidPhone(c.Phone) {
		return Contact{}, errors.New("Phone must be in E.164 format (e.g. +1234567890)")
	}

	if c.Name == "Anton" {
		return Contact{}, errors.New("name cannot be Anton")
	}

	return s.repo.Update(id, c)
}

func (s *contactService) DeleteContactByID(id int) error {
	return s.repo.DeleteByID(id)
}

func (s *contactService) DeleteContactByName(name string) error {
	return s.repo.DeleteByName(name)
}