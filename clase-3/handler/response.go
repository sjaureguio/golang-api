package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Success = true
	Error   = false
	Message = "message"
)

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(success bool, message string, data interface{}) response {
	return response{
		success,
		message,
		data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
