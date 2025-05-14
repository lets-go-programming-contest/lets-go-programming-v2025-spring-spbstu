package grpchandler

import (
	"context"
	"errors"

	grpcserver "github.com/quaiion/go-practice/grpc-contact-manager/gen/proto/contact_manager/v1"
	"github.com/quaiion/go-practice/grpc-contact-manager/internal/cm"

	"github.com/bufbuild/protovalidate-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	grpcserver.UnimplementedContactManagerServiceServer
	manager   *cm.ContMan
	validator protovalidate.Validator
}

func New(manager *cm.ContMan, validator protovalidate.Validator) *Handler {
	return &Handler{
		manager:   manager,
		validator: validator,
	}
}

func (h *Handler) AddContact(ctx context.Context, req *grpcserver.AddContactRequest) (*grpcserver.AddContactResponse, error) {
        err := h.validator.Validate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

        contact := cm.NewContact(req.Name, req.Number)

	err = h.manager.Add(contact)
	if err != nil {
		if errors.Is(err, cm.ErrNameRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, cm.ErrDuplicateAdded) {
                        return nil, status.Error(codes.AlreadyExists, err.Error())
                } else {
		        return nil, status.Error(codes.Internal, err.Error())
                }
	}

	return &grpcserver.AddContactResponse{ Contact: convertContact(&contact) }, nil
}

func (h *Handler) GetContact(ctx context.Context, req *grpcserver.GetContactRequest) (*grpcserver.GetContactResponse, error) {
	err := h.validator.Validate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contact, err := h.manager.Get(req.Id)
	if err != nil {
                if errors.Is(err, cm.ErrGetContNotFound) {
                        return nil, status.Error(codes.NotFound, err.Error())
                } else {
                        return nil, status.Error(codes.Internal, err.Error())
                }
	}

	return &grpcserver.GetContactResponse{ Contact: convertContact(&contact) }, nil
}

func (h *Handler) GetAllContacts(ctx context.Context, req *grpcserver.GetAllContactsRequest) (*grpcserver.GetAllContactsResponse, error) {
	err := h.validator.Validate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contacts, err := h.manager.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoContacts := make([]*grpcserver.Contact, len(contacts))
	for i, contact := range contacts {
		protoContacts[i] = convertContact(&contact)
	}

	return &grpcserver.GetAllContactsResponse{ Contacts: protoContacts }, nil
}

func (h *Handler) UpdateContact(ctx context.Context, req *grpcserver.UpdateContactRequest) (*grpcserver.UpdateContactResponse, error) {
	err := h.validator.Validate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

        contact := cm.NewContact(req.Name, req.Number)

	err = h.manager.Update(contact)
	if err != nil {
		if errors.Is(err, cm.ErrDuplicateAdded) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, cm.ErrNameRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, cm.ErrContUpdNotFound) {
                        return nil, status.Error(codes.NotFound, err.Error())
                } else {
		        return nil, status.Error(codes.Internal, err.Error())
                }
	}

	return &grpcserver.UpdateContactResponse{ Contact: convertContact(&contact) }, nil
}

func (h *Handler) DeleteContact(ctx context.Context, req *grpcserver.DeleteContactRequest) (*grpcserver.DeleteContactResponse, error) {
	err := h.validator.Validate(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.manager.Delete(req.Id)
	if err != nil {
		if errors.Is(err, cm.ErrContDelNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
                        return nil, status.Error(codes.Internal, err.Error())
                }
	}

	return &grpcserver.DeleteContactResponse{}, nil
}

func convertContact(contact *cm.Contact) *grpcserver.Contact {
	return &grpcserver.Contact{
		Id:     contact.ID,
		Name:   contact.Name,
		Number: contact.Number,
	}
}
