package httputil

import (
	"encoding/json"
	"net/http"
)

// JSON sends data in JSON format with the specified HTTP status
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// NoContent sends a 204 No Content response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
