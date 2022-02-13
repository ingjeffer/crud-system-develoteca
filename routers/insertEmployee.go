package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ingjeffer/crud-system-develoteca/db"
	"github.com/ingjeffer/crud-system-develoteca/models"
)

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee

	reqBody, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &employee)
	// err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "has a error "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(employee)

	status, err2, id := db.InsertEmployee(employee)
	if err2 != nil {
		http.Error(w, "has a error2 "+err2.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "has a error into insert employee ", http.StatusBadRequest)
		return
	}
	employee.Id = id
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}
