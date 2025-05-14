package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/dmitriy.rumyantsev/task-9/internal/domain"
	"github.com/dmitriy.rumyantsev/task-9/internal/service"
	"github.com/dmitriy.rumyantsev/task-9/pkg/httputil"
)

// ContactHandler processes requests for contact management
type ContactHandler struct {
	contactService service.ContactService
}

// NewContactHandler creates a new contact handler
func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

// Register registers route handlers
func (h *ContactHandler) Register(mux *http.ServeMux) {
	// Handle requests to /contacts
	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetContacts(w, r)
		case http.MethodPost:
			h.CreateContact(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Handle requests to /contacts/{id}
	idRegex := regexp.MustCompile(`^/contacts/(\d+)$`)
	mux.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		matches := idRegex.FindStringSubmatch(r.URL.Path)
		if len(matches) < 2 {
			httputil.JSONError(w, httputil.NewBadRequestError("Invalid URL", "Expected format: /contacts/{id}"))
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			httputil.JSONError(w, httputil.NewBadRequestError("Invalid ID", "ID must be a number"))
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetContact(w, r, id)
		case http.MethodPut:
			h.UpdateContact(w, r, id)
		case http.MethodDelete:
			h.DeleteContact(w, r, id)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// GetContacts handles GET /contacts
func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contacts, err := h.contactService.GetAll(ctx)
	if err != nil {
		var errResp httputil.ErrorResponse
		if errors.As(err, &errResp) {
			httputil.JSONError(w, errResp)
		} else {
			httputil.JSONError(w, httputil.NewInternalServerError(err.Error()))
		}
		return
	}

	httputil.JSON(w, http.StatusOK, contacts)
}

// GetContact handles GET /contacts/{id}
func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request, id int) {
	ctx := r.Context()

	contact, err := h.contactService.GetByID(ctx, id)
	if err != nil {
		var errResp httputil.ErrorResponse
		if errors.As(err, &errResp) {
			httputil.JSONError(w, errResp)
		} else {
			httputil.JSONError(w, httputil.NewInternalServerError(err.Error()))
		}
		return
	}

	httputil.JSON(w, http.StatusOK, contact)
}

// CreateContact handles POST /contacts
func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		httputil.JSONError(w, httputil.NewBadRequestError("Invalid request body", err.Error()))
		return
	}

	ctx := r.Context()

	newContact, err := h.contactService.Create(ctx, contact)
	if err != nil {
		var errResp httputil.ErrorResponse
		if errors.As(err, &errResp) {
			httputil.JSONError(w, errResp)
		} else {
			httputil.JSONError(w, httputil.NewInternalServerError(err.Error()))
		}
		return
	}

	httputil.JSON(w, http.StatusCreated, newContact)
}

// UpdateContact handles PUT /contacts/{id}
func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request, id int) {
	var contact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		httputil.JSONError(w, httputil.NewBadRequestError("Invalid request body", err.Error()))
		return
	}

	ctx := r.Context()

	updatedContact, err := h.contactService.Update(ctx, id, contact)
	if err != nil {
		var errResp httputil.ErrorResponse
		if errors.As(err, &errResp) {
			httputil.JSONError(w, errResp)
		} else {
			httputil.JSONError(w, httputil.NewInternalServerError(err.Error()))
		}
		return
	}

	httputil.JSON(w, http.StatusOK, updatedContact)
}

// DeleteContact handles DELETE /contacts/{id}
func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request, id int) {
	ctx := r.Context()

	err := h.contactService.Delete(ctx, id)
	if err != nil {
		var errResp httputil.ErrorResponse
		if errors.As(err, &errResp) {
			httputil.JSONError(w, errResp)
		} else {
			httputil.JSONError(w, httputil.NewInternalServerError(err.Error()))
		}
		return
	}

	httputil.NoContent(w)
}
