package model

import (
	"time"

	contactpb "task-10/gen/proto/api/proto/contact_manager/v1"
)

type Contact struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (c *Contact) ToProto() *contactpb.Contact {
	return &contactpb.Contact{
		Id:        c.ID,
		Name:      c.Name,
		Phone:     c.Phone,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Unix(),
		UpdatedAt: c.UpdatedAt.Unix(),
	}
}

func ContactFromProto(pb *contactpb.Contact) *Contact {
	return &Contact{
		ID:        pb.Id,
		Name:      pb.Name,
		Phone:     pb.Phone,
		Email:     pb.Email,
		CreatedAt: time.Unix(pb.CreatedAt, 0),
		UpdatedAt: time.Unix(pb.UpdatedAt, 0),
	}
}
