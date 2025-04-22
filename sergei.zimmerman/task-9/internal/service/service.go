package service

import (
	"github.com/google/uuid"
	"phonebook/internal/errors"
	"phonebook/internal/phonebook"
)

type Service struct {
	repo phonebook.Repository
}

func New(repo phonebook.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll() ([]phonebook.Contact, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uuid.UUID) (phonebook.Contact, error) {
	contact, err := s.repo.GetByID(id)
	if err != nil {
		return contact, errors.ErrNotFound
	}

	return contact, nil
}

func (s *Service) Create(name, phone string) (phonebook.Contact, error) {
	c := phonebook.Contact{
		ID:    uuid.New(),
		Name:  name,
		Phone: phone,
	}
	err := s.repo.Create(c)

	return c, err
}

func (s *Service) Update(id uuid.UUID, name, phone string) error {
	return s.repo.Update(id, phonebook.Contact{Name: name, Phone: phone})
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
