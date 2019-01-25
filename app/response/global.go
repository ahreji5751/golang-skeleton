package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GlobalResponse struct {
}

func (gr GlobalResponse) WithError(w http.ResponseWriter, errorCode int, status string, message string) {
	error := ErrorResponse{status, message}
	gr.WithJson(w, errorCode, error)
}

func (gr GlobalResponse) WithJson(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
