package http

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

// Returns Success Response as JSON
// Use for handlers
func NewSuccessResponse(w http.ResponseWriter, statusCode int, response string){
	success := SuccessResponse{
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&success)
	return
}