package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ingjeffer/crud-system-develoteca/constants"
)

func ExecuteResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set(constants.ContentType, constants.ApplicationJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
