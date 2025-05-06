package service

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"

	"task-9/internal/phonebook"
)

type Service struct {
	pb phonebook.IPhonebook
}

func New(pb phonebook.IPhonebook) *Service {
	return &Service{pb}
}

func (s *Service) GetAll() ([]phonebook.Contact, error) {
	return s.pb.GetAll()
}

func (s *Service) GetByID(id uuid.UUID) (phonebook.Contact, error) {
	contact, err := s.pb.GetByID(id)
	if err != nil {
		return contact, fmt.Errorf("not found by id %v", err)
	}
	return contact, nil
}

func validatePhone(phone string) error {
	re := regexp.MustCompile(`\+[0-9]*`)
	if !re.MatchString(phone) {
		return fmt.Errorf("invalid phone format")
	}
	return nil
}

func (s *Service) Create(name string, phone string) (phonebook.Contact, error) {
	err := validatePhone(phone)
	if err != nil {
		return phonebook.Contact{}, fmt.Errorf("fail to create contact: %v", err)
	}
	c := phonebook.Contact{
		ID:    uuid.New(),
		Name:  name,
		Phone: phone,
	}
	err = s.pb.Create(c)
	return c, err
}

func (s *Service) Update(id uuid.UUID, name string, phone string) error {
	return s.pb.Update(id, phonebook.Contact{Name: name, Phone: phone})
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.pb.Delete(id)
}
