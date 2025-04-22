package manager

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrContactNotFound = errors.New("contact not found")
	ErrInvalidInput    = errors.New("invalid input")
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) CreateContact(ctx context.Context, name, phone string) (*Contact, error) {
	if name == "" || phone == "" {
		return nil, ErrInvalidInput
	}

	contact := &Contact{
		ID:    uuid.New().String(),
		Name:  name,
		Phone: phone,
	}

	if err := m.db.WithContext(ctx).Create(contact).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (m *Manager) GetContact(ctx context.Context, id string) (*Contact, error) {
	var contact Contact
	if err := m.db.WithContext(ctx).First(&contact, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return &contact, nil
}

func (m *Manager) ListContacts(ctx context.Context) ([]*Contact, error) {
	var contacts []*Contact
	if err := m.db.WithContext(ctx).Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (m *Manager) UpdateContact(ctx context.Context, id, name, phone string) (*Contact, error) {
	if name == "" || phone == "" {
		return nil, ErrInvalidInput
	}

	contact := &Contact{
		ID:    id,
		Name:  name,
		Phone: phone,
	}

	if err := m.db.WithContext(ctx).First(contact, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}

	if err := m.db.WithContext(ctx).Model(contact).Updates(map[string]interface{}{
		"name":  name,
		"phone": phone,
	}).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (m *Manager) DeleteContact(ctx context.Context, id string) error {
	result := m.db.WithContext(ctx).Delete(&Contact{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound
	}
	return nil
}
