package utils

import (
	"encoding/json"
	"net/http"
)

type successRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type errorRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data any) {
	res := successRes{
		Success: true,
		Message: message,
		Data:    data,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	res := errorRes{
		Success: false,
		Message: message,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
