package handler

import (
	"context"
	"errors"
	contactpb "task-10/gen/proto/api/proto/contact_manager/v1"
	"task-10/internal/model"
	"task-10/internal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContactServiceHandler struct {
	contactpb.UnimplementedContactServiceServer
	store *storage.ContactStore
}

func NewContactServiceHandler(store *storage.ContactStore) *ContactServiceHandler {
	return &ContactServiceHandler{store: store}
}

func (h *ContactServiceHandler) CreateContact(ctx context.Context, req *contactpb.CreateContactRequest) (*contactpb.ContactResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contact := model.Contact{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	}

	createdContact, err := h.store.Create(contact)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrDuplicateContact):
			return nil, status.Error(codes.AlreadyExists, "contact already exists")
		case errors.Is(err, storage.ErrInvalidData):
			return nil, status.Error(codes.InvalidArgument, "invalid data provided")
		default:
			return nil, status.Error(codes.Internal, "failed to create contact")
		}
	}

	return &contactpb.ContactResponse{Contact: createdContact.ToProto()}, nil
}

func (h *ContactServiceHandler) GetContact(ctx context.Context, req *contactpb.GetContactRequest) (*contactpb.ContactResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contact, err := h.store.GetByID(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "contact not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contact")
	}

	return &contactpb.ContactResponse{Contact: contact.ToProto()}, nil
}

func (h *ContactServiceHandler) ListContacts(ctx context.Context, req *contactpb.ListContactsRequest) (*contactpb.ListContactsResponse, error) {
	contacts, err := h.store.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list contacts")
	}

	pbContacts := make([]*contactpb.Contact, len(contacts))
	for i, contact := range contacts {
		pbContacts[i] = contact.ToProto()
	}

	return &contactpb.ListContactsResponse{Contacts: pbContacts}, nil
}

func (h *ContactServiceHandler) UpdateContact(ctx context.Context, req *contactpb.UpdateContactRequest) (*contactpb.ContactResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contact := model.Contact{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	}

	updatedContact, err := h.store.Update(req.Id, contact)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return nil, status.Error(codes.NotFound, "contact not found")
		case errors.Is(err, storage.ErrDuplicateContact):
			return nil, status.Error(codes.AlreadyExists, "contact with this phone already exists")
		case errors.Is(err, storage.ErrInvalidData):
			return nil, status.Error(codes.InvalidArgument, "invalid data provided")
		default:
			return nil, status.Error(codes.Internal, "failed to update contact")
		}
	}

	return &contactpb.ContactResponse{Contact: updatedContact.ToProto()}, nil
}

func (h *ContactServiceHandler) DeleteContact(ctx context.Context, req *contactpb.GetContactRequest) (*contactpb.DeleteContactResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := h.store.Delete(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "contact not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete contact")
	}

	return &contactpb.DeleteContactResponse{Message: "Contact successfully deleted"}, nil
}
