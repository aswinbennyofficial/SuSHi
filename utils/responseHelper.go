package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aswinbennyofficial/SuSHi/models"
)


func ResponseHelper(w http.ResponseWriter, code int, message string, data interface{}){
	response := models.Response{
		Status: http.StatusText(code),
		Message: message,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}