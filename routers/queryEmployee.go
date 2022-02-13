package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ingjeffer/crud-system-develoteca/db"
	"github.com/ingjeffer/crud-system-develoteca/usermodels"
	"github.com/ingjeffer/crud-system-develoteca/utils"
)

func QueryEmployee(w http.ResponseWriter, r *http.Request) {
	status, err, registers := db.QueryEmployee()
	if err != nil {
		utils.ExecuteResponse(w, "Has error "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		utils.ExecuteResponse(w, "Has error with query status ", http.StatusBadRequest)
		return
	}

	var response usermodels.EmployeesResponse
	response.Data = registers
	response.Success = true
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
