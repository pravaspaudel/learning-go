package utils

import (
	"encoding/json"
	"net/http"

	"github.com/pravaspaudel/09_crud/internals/models"
)

func SendSuccessResponse(w http.ResponseWriter, msg string, status int, data any) {
	response := models.SuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, msg string, status int) {
	response := models.ErrorResponse{
		Success: false,
		Message: msg,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
