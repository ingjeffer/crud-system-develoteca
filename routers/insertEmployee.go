package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ingjeffer/crud-system-develoteca/db"
	"github.com/ingjeffer/crud-system-develoteca/models"
	"github.com/ingjeffer/crud-system-develoteca/usermodels"
	"github.com/ingjeffer/crud-system-develoteca/utils"
)

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	var responseDefault usermodels.ResponseDefault

	reqBody, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &employee)
	// err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		// http.Error(w, "has a error "+err.Error(), http.StatusBadRequest)
		responseDefault.Success = false
		responseDefault.ErrorMessage = err.Error()
		utils.ExecuteResponse(w, responseDefault, http.StatusBadRequest)
		return
	}

	fmt.Println(employee)

	status, err2, id := db.InsertEmployee(employee)
	if err2 != nil {
		// http.Error(w, "has a error2 "+err2.Error(), http.StatusBadRequest)
		responseDefault.Success = false
		responseDefault.ErrorMessage = err2.Error()
		utils.ExecuteResponse(w, responseDefault, http.StatusBadRequest)
		return
	}
	if !status {
		// http.Error(w, "has a error into insert employee ", http.StatusBadRequest)
		responseDefault.Success = false
		responseDefault.ErrorMessage = "has a error into insert employee "
		utils.ExecuteResponse(w, responseDefault, http.StatusBadRequest)
		return
	}
	employee.Id = id
	var response usermodels.EmployeeResponse
	response.Data = employee
	response.Success = true
	// utils.ExecuteResponse(w, response, http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
