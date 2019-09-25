package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

// Returns Error Response as JSON
// Use for handlers
func NewErrorResponse(w http.ResponseWriter, statusCode int, response string){
	error := ErrorResponse{
		true,
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}