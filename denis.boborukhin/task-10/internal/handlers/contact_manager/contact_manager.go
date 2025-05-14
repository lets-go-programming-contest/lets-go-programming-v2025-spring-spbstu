package contact_manager

import (
	"context"

	v1 "github.com/denisboborukhin/contact_manager/gen/proto/contact_manager/v1"
	"github.com/denisboborukhin/contact_manager/internal/contact/manager"

	"github.com/bufbuild/protovalidate-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	v1.UnimplementedContactManagerServiceServer
	manager   *manager.Manager
	validator protovalidate.Validator
}

func NewHandler(manager *manager.Manager, validator protovalidate.Validator) *Handler {
	return &Handler{
		manager:   manager,
		validator: validator,
	}
}

func (h *Handler) CreateContact(ctx context.Context, req *v1.CreateContactRequest) (*v1.CreateContactResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	contact, err := h.manager.CreateContact(ctx, req.Name, req.Phone)
	if err != nil {
		if err == manager.ErrInvalidInput {
			return nil, status.Error(codes.InvalidArgument, "invalid input: name and phone are required")
		}
		return nil, status.Error(codes.Internal, "failed to create contact")
	}

	return &v1.CreateContactResponse{
		Contact: convertContact(contact),
	}, nil
}

func (h *Handler) GetContact(ctx context.Context, req *v1.GetContactRequest) (*v1.GetContactResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	contact, err := h.manager.GetContact(ctx, req.Id)
	if err != nil {
		if err == manager.ErrContactNotFound {
			return nil, status.Error(codes.NotFound, "contact not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contact")
	}

	return &v1.GetContactResponse{
		Contact: convertContact(contact),
	}, nil
}

func (h *Handler) ListContacts(ctx context.Context, req *v1.ListContactsRequest) (*v1.ListContactsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	contacts, err := h.manager.ListContacts(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list contacts")
	}

	protoContacts := make([]*v1.Contact, len(contacts))
	for i, contact := range contacts {
		protoContacts[i] = convertContact(contact)
	}

	return &v1.ListContactsResponse{
		Contacts: protoContacts,
	}, nil
}

func (h *Handler) UpdateContact(ctx context.Context, req *v1.UpdateContactRequest) (*v1.UpdateContactResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	contact, err := h.manager.UpdateContact(ctx, req.Id, req.Name, req.Phone)
	if err != nil {
		if err == manager.ErrContactNotFound {
			return nil, status.Error(codes.NotFound, "contact not found")
		}
		if err == manager.ErrInvalidInput {
			return nil, status.Error(codes.InvalidArgument, "invalid input: name and phone are required")
		}
		return nil, status.Error(codes.Internal, "failed to update contact")
	}

	return &v1.UpdateContactResponse{
		Contact: convertContact(contact),
	}, nil
}

func (h *Handler) DeleteContact(ctx context.Context, req *v1.DeleteContactRequest) (*v1.DeleteContactResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	err := h.manager.DeleteContact(ctx, req.Id)
	if err != nil {
		if err == manager.ErrContactNotFound {
			return nil, status.Error(codes.NotFound, "contact not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete contact")
	}

	return &v1.DeleteContactResponse{}, nil
}

func convertContact(contact *manager.Contact) *v1.Contact {
	return &v1.Contact{
		Id:    contact.ID,
		Name:  contact.Name,
		Phone: contact.Phone,
	}
}
