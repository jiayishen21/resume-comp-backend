package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(&payload)
}

func WriteJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, map[string]string{"error": err.Error()})
}
