package httperr

import (
	"encoding/json"
	"errors"
	"net/http"
)

var errFailedEncode = errors.New("failed to encode error message")

type HttpErr struct {
        Msg string `json:"error_message"`
}

func New(err error) HttpErr {
        return HttpErr{ Msg: err.Error() }
}

func (httpErr HttpErr) Encode(writer http.ResponseWriter, status int) {
        errMsg, err := json.Marshal(httpErr)
        if err != nil {
                status = http.StatusInternalServerError
                errMsg, err = json.Marshal(New(errFailedEncode))
                if err != nil {
                        panic(errFailedEncode)
                }
        }

        http.Error(writer, string(errMsg), status)
}
