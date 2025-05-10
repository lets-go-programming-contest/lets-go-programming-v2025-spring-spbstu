package resthandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/quaiion/go-practice/grpc-contact-manager/internal/cm"
)

type ContactDB interface {
        GetAll() ([]cm.Contact, error)
	Get(string) (cm.Contact, error)
	Add(cm.Contact) error
	Update(cm.Contact) error
	Delete(string) error
}

type Handler struct {
	cdb ContactDB
}

func New(cdb ContactDB) *Handler {
	return &Handler{ cdb: cdb }
}

var (
        errGetAllFailed     = errors.New("'get all' request processing failed")
        errGetFailed        = errors.New("'get' request processing failed")
        errUpdFailed        = errors.New("'update' request processing failed")
        errDelFailed        = errors.New("'delete' request processing failed")
        errAddFailed        = errors.New("'add' request processing failed")
        errDecodeFailed     = errors.New("failed decoding the input contact")
        errRegexInitFailed  = errors.New("failed to initialize number format checker")
        errFormatViolated   = errors.New("number format violated")
        errEmptyID          = errors.New("input ID should not be empty")
        errMethodNotAllowed = errors.New("method not allowed")
)

func (h *Handler) HandleAllContacts(writer http.ResponseWriter, request *http.Request) {
        switch request.Method {
        case http.MethodGet:
                contacts, err := h.cdb.GetAll()
                if err != nil {
                        genErrMsg(writer, errors.Join(errGetAllFailed, err), http.StatusInternalServerError)
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                json.NewEncoder(writer).Encode(contacts)

        case http.MethodPost:
                var contact cm.Contact

                err := json.NewDecoder(request.Body).Decode(&contact)
                if err != nil {
                        genErrMsg(writer, errors.Join(errDecodeFailed, err), http.StatusBadRequest)
                        return
                }

                numberRegExp, err := regexp.Compile(`^\\+[1-9][0-9]?[0-9]?[0-9]{11}$`)
                if err != nil {
                        genErrMsg(writer, errors.Join(errRegexInitFailed, err), http.StatusInternalServerError)
                        return
                }

                if !numberRegExp.MatchString(contact.Number) {
                        genErrMsg(writer, errFormatViolated, http.StatusBadRequest)
                        return
                }

                err = h.cdb.Add(contact)
                if err != nil {
                        if errors.Is(err, cm.ErrNameRequired) {
                                genErrMsg(writer, err, http.StatusBadRequest)
                        } else if errors.Is(err, cm.ErrDuplicateAdded) {
                                genErrMsg(writer, errors.Join(errAddFailed, err), http.StatusConflict)
                        } else {
                                genErrMsg(writer, errors.Join(errAddFailed, err), http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusCreated)

        default:
                genErrMsg(writer, errMethodNotAllowed, http.StatusMethodNotAllowed)
        }
}

func (h *Handler) HandleContact(writer http.ResponseWriter, request *http.Request) {
        id := request.URL.Path[len(`/contacts/`):]
        if id == `` {
                genErrMsg(writer, errEmptyID, http.StatusBadRequest)
                return
        }

        switch request.Method {

        case http.MethodGet:
                contact, err := h.cdb.Get(id)
                if err != nil {
                        if errors.Is(err, cm.ErrGetContNotFound) {
                                genErrMsg(writer, errors.Join(errGetFailed, err), http.StatusNotFound)
                        } else {
                                genErrMsg(writer, errors.Join(errGetFailed, err), http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                json.NewEncoder(writer).Encode(contact)

        case http.MethodPut:
                var contact cm.Contact

                err := json.NewDecoder(request.Body).Decode(&contact)
                if err != nil {
                        genErrMsg(writer, errors.Join(errDecodeFailed, err), http.StatusBadRequest)
                        return
                }

                numberRegExp, err := regexp.Compile(`^\\+[1-9][0-9]?[0-9]?[0-9]{11}$`)
                if err != nil {
                        genErrMsg(writer, errors.Join(errRegexInitFailed, err), http.StatusInternalServerError)
                        return
                }

                if !numberRegExp.MatchString(contact.Number) {
                        genErrMsg(writer, errFormatViolated, http.StatusBadRequest)
                        return
                }

                contact.ID = id
                err = h.cdb.Update(contact)
                if err != nil {
                        if errors.Is(err, cm.ErrNameRequired) {
                                genErrMsg(writer, err, http.StatusBadRequest)
                        } else if errors.Is(err, cm.ErrDuplicateAdded) {
                                genErrMsg(writer, errors.Join(errUpdFailed, err), http.StatusConflict)
                        } else if errors.Is(err, cm.ErrContUpdNotFound) {
                                genErrMsg(writer, errors.Join(errUpdFailed, err), http.StatusNotFound)
                        } else {
                                genErrMsg(writer, errors.Join(errUpdFailed, err), http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusOK)

        case http.MethodDelete:
                err := h.cdb.Delete(id)
                if err != nil {
                        if errors.Is(err, cm.ErrContDelNotFound) {
                                genErrMsg(writer, errors.Join(errDelFailed, err), http.StatusNotFound)
                        } else {
                                genErrMsg(writer, errors.Join(errDelFailed, err), http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusOK)

        default:
                genErrMsg(writer, errMethodNotAllowed, http.StatusMethodNotAllowed)
        }
}

func genErrMsg(writer http.ResponseWriter, err error, status int) {
        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, status)
}
