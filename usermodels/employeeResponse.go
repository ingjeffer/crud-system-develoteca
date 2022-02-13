package usermodels

import "github.com/ingjeffer/crud-system-develoteca/models"

type EmployeeResponse struct {
	ResponseDefault
	Data models.Employee `json:"data,omitempty"`
}

type EmployeesResponse struct {
	ResponseDefault
	Data []models.Employee `json:"data,omitempty"`
}
