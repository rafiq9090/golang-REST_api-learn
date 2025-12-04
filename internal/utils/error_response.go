package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

func JSONError(w http.ResponseWriter, massage string, statusCode int, details map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Error:   massage,
		Details: details,
	}
	json.NewEncoder(w).Encode(response)
}
